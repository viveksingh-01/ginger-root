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

	var userID, guestID string

	if val, exists := c.Get("userId"); exists {
		userID = val.(string)
		guestID = ""
	} else {
		guestID = c.GetHeader("X-Guest-Id")
		if guestID == "" {
			guestID = generateGuestID()
			c.Header("X-Guest-Id", guestID)
		}
	}

	resp, err := h.service.AddToCart(
		c.Request.Context(),
		userID,
		guestID,
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

func (h *Handler) GetCart(c *gin.Context) {
	var userID, guestID string

	if val, exists := c.Get("userId"); exists {
		userID = val.(string)
	} else {
		guestID = c.GetHeader("X-Guest-Id")
		if guestID == "" {
			c.JSON(http.StatusOK, gin.H{
				"statusCode":    0,
				"statusMessage": "SUCCESS",
				"data": gin.H{
					"cartMeta":  gin.H{},
					"cartItems": gin.H{"items": []any{}},
				},
			})
			return
		}
	}

	resp, err := h.service.GetCart(
		c.Request.Context(),
		userID,
		guestID,
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
