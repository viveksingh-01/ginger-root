package order

import (
	"github.com/viveksingh-01/ginger-root/internal/address"
	"github.com/viveksingh-01/ginger-root/internal/cart"
)

type OrderRequest struct {
	CartID        string `json:"cartId" binding:"required"`
	AddressID     string `json:"addressId" binding:"required"`
	PaymentMethod string `json:"paymentMethod" binding:"required"`
}

type OrderResponse struct {
	ID             string                  `json:"id"`
	OrderID        int                     `json:"orderId"`
	RestaurantID   string                  `json:"restaurantId"`
	RestaurantName string                  `json:"restaurantName"`
	Address        address.AddressResponse `json:"address"`
	Items          []cart.Item             `json:"items"`
	BillDetails    BillDetails             `json:"billDetails"`
	Status         string                  `json:"status"`
}

type BillDetails struct {
	Subtotal       float64 `json:"subtotal"`
	DeliveryCharge float64 `json:"deliveryCharge"`
	GST            float64 `json:"gst"`
	Discount       float64 `json:"discount"`
	FinalAmount    float64 `json:"finalAmount"`
}

type Response struct {
	StatusCode    int            `json:"statusCode"`
	StatusMessage string         `json:"statusMessage"`
	Data          *OrderResponse `json:"data,omitempty"`
}

type OrderStatusEvent struct {
	OrderID        int    `json:"orderId"`
	RestaurantName string `json:"restaurantName"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle,omitempty"`
	Step           int    `json:"step"`
	TotalSteps     int    `json:"totalSteps"`
	ETA            int    `json:"eta"`
	IsTerminal     bool   `json:"isTerminal"`
	UpdatedAt      string `json:"updatedAt"`
}
