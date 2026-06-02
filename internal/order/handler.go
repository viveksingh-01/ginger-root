package order

import (
	"context"
	"errors"
	"net/http"
	"strconv"

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

func (h *Handler) TrackOrderStatus(c *gin.Context) {
	orderID, err := parseOrderID(c.Param("orderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode":    1,
			"statusMessage": "invalid orderId; expected a 6-digit number",
		})
		return
	}

	userID := c.GetString("userId")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"statusCode":    1,
			"statusMessage": "unauthorized",
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	streamErr := h.service.StreamOrderStatus(c.Request.Context(), userID, orderID, func(ev OrderStatusEvent) error {
		c.SSEvent("order-status", ev)
		c.Writer.Flush()
		return nil
	})

	if streamErr == nil {
		return
	}

	if errors.Is(streamErr, context.Canceled) || errors.Is(streamErr, context.DeadlineExceeded) {
		return
	}

	if c.Writer.Written() {
		return
	}
}

func parseOrderID(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < 100000 || n > 999999 {
		return 0, errors.New("invalid orderId")
	}
	return n, nil
}

func toString(v any) string {
	if v == nil {
		return ""
	}
	return v.(string)
}
