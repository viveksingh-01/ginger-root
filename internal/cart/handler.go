package cart

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

func (h *Handler) AddToCart(c *gin.Context) {
	var req AddToCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	userIDVal, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			StatusCode:    1,
			StatusMessage: "unauthorized",
		})
		return
	}

	userID := userIDVal.(string)

	resp, err := h.service.AddToCart(
		c.Request.Context(),
		userID,
		req.Cart.RestaurantID,
		req.Cart.AddressID,
		req.Cart.CartItems,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode:    0,
		StatusMessage: "fetched successfully",
		Data:          resp.Data,
	})
}
