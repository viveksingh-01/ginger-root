package order

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

func (h *Handler) PlaceOrder(c *gin.Context) {
	var req OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	userID, _ := c.Get("userId")

	order, err := h.service.PlaceOrder(
		c.Request.Context(),
		toString(userID),
		req.CartID,
		req.AddressID,
		req.PaymentMethod,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode:    0,
		StatusMessage: "fetched successfully",
		Data:          ToOrderResponse(order),
	})
}

func toString(v any) string {
	if v == nil {
		return ""
	}
	return v.(string)
}
