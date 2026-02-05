package menu

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

func (h *Handler) GetMenu(c *gin.Context) {
	restaurantID := c.Param("id")
	menu, err := h.service.GetMenu(c.Request.Context(), restaurantID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "menu not found",
		})
	}
	c.JSON(http.StatusOK, menu)
}
