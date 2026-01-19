package restaurant

import (
	"context"
	"log"
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
	// Read limit and offset query params from the request and parse them to int64 type
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "20"), 10, 64)
	offset, _ := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)

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
	restaurants, err := h.service.List(ctx, limit, offset)
	if err != nil {
		log.Println("Error occurred while fetching restaurants", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch restaurants"})
		return
	}

	// Sends HTTP 200 OK
	// Serializes restaurants to JSON
	// Gin automatically:
	// * Sets Content-Type: application/json
	// * Marshals Go structs into JSON
	c.JSON(http.StatusOK, gin.H{
		"count":  len(restaurants),
		"data":   restaurants,
		"limit":  limit,
		"offset": offset,
	})
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
