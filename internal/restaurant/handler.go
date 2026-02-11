package restaurant

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// This code:
// * Handles incoming HTTP requests
// * Calls the service layer to perform business logic
// * Converts results into HTTP responses (JSON)
// * Handles errors and HTTP status codes
// This is the outermost layer of our application.

// Handler depends on the Service
// It delegates all business work to the service
// It only handles HTTP concerns
type Handler struct {
	service *Service
}

// Creates a new handler
// Injects the service dependency
// Keeps handler easy to test and initialize
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// This method is an HTTP endpoint (registered to a route)
func (h *Handler) List(c *gin.Context) {
	// Read page and limit query params from the request
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		h.badRequest(c, "INVALID_PAGE_VALUE", "page must be greater than 0")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		h.badRequest(c, "INVALID_LIMIT", "limit must be between 1 and 100")
		return
	}

	skip := (page - 1) * limit

	// Filter veg-only restaurants
	var veg *bool
	if vegStr := c.Query("veg"); vegStr != "" {
		v, err := strconv.ParseBool(vegStr)
		if err != nil {
			h.badRequest(c, "INVALID_VEG_FILTER_VALUE", "veg must be true or false")
			return
		}
		veg = &v
	}

	// Filter by minimum rating
	var minRating *float64
	if ratingStr := c.Query("minRating"); ratingStr != "" {
		r, err := strconv.ParseFloat(ratingStr, 64)
		if err != nil || r < 0 || r > 5 {
			h.badRequest(c, "INVALID_RATING_VALUE", "minRating must be between 0 and 5")
			return
		}
		minRating = &r
	}

	filter := Filter{
		Veg:       veg,
		MinRating: minRating,
	}

	// Apply sorting
	sortBy := c.DefaultQuery("sortBy", "rating")
	order := c.DefaultQuery("order", "desc")

	// Convert client-facing sort parameter into a safe MongoDB field name
	allowedSorts := map[string]string{
		"rating":   "avg_rating",
		"delivery": "sla.delivery_time",
	}
	sortField := allowedSorts[sortBy]

	// Set sort-order based on the input
	sortOrder := -1
	if order == "asc" {
		sortOrder = 1
	}

	sort := Sort{
		Field: sortField,
		Order: sortOrder,
	}

	// Create a context that:
	// 1. Automatically cancels after 3 seconds
	// 2. Signals MongoDB to stop work
	// 3. Frees resources

	// c.Request.Context() extracts the standard Go context.Context, it:
	// 1. Is cancelled if the client disconnects
	// 2. Carries request-scoped deadlines
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	// The handler calls the service to fetch list of restaurants
	restaurants, err := h.service.List(ctx, int64(skip), int64(limit), filter, sort)
	if err != nil {
		h.internalError(c, err)
		return
	}

	// Sends HTTP 200 OK
	// Serializes restaurants to JSON
	// Gin automatically:
	// * Sets Content-Type: application/json
	// * Marshals Go structs into JSON
	c.JSON(http.StatusOK, gin.H{
		"count": len(restaurants),
		"data":  restaurants,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetRestaurant(c *gin.Context) {
	restaurantID := c.Param("restaurantId")
	restaurant, err := h.service.GetRestaurant(c.Request.Context(), restaurantID)
	if err != nil {
		h.internalError(c, err)
		return
	}
	c.JSON(http.StatusOK, restaurant)
}

// IMPORTANT:

// Context timeouts — WHY THIS IS CRITICAL
// What is a context?
// A context:
// * Carries deadlines
// * Can be cancelled
// * Propagates across layers

// What happens without this?
// 1. Mongo query hangs
// 2. Goroutine never exits
// 3. App slows down under load

// Why do we pass ctx everywhere? - h.service.List(ctx, limit, offset)
// Because:
// 1. Handler controls request lifetime
// 2. Service & repository obey it
// 3. Cancellation propagates downward
// This is idiomatic Go.
