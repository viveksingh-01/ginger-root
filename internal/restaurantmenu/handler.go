package restaurantmenu

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetRestaurantMenu(c *gin.Context) {
	restaurantID := c.Param("restaurantId")
	r, err := h.service.GetRestaurantMenu(c.Request.Context(), restaurantID)
	if err != nil {
		if errors.Is(err, restaurant.ErrRestaurantNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}
