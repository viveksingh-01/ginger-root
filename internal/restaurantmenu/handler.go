package restaurantmenu

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetRestaurantMenu(c *gin.Context) {
	restaurantID := c.Param("restaurantId")
	restaurantMenu, err := h.service.GetRestaurantMenu(c, restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, restaurantMenu)
}
