package order

import (
	"time"

	"github.com/viveksingh-01/ginger-root/internal/address"
)

type Order struct {
	ID             string          `bson:"id" json:"id"`
	UserID         string          `bson:"userId,omitempty"`
	GuestID        string          `bson:"guestId,omitempty"`
	RestaurantID   string          `bson:"restaurantId"`
	RestaurantName string          `bson:"restaurantName"`
	Address        address.Address `bson:"address"`
	Items          []OrderItem     `bson:"items"`
	Subtotal       float64         `bson:"subtotal"`
	Delivery       float64         `bson:"delivery"`
	GST            float64         `bson:"gst"`
	Discount       float64         `bson:"discount"`
	FinalAmount    float64         `bson:"finalAmount"`
	PaymentMethod  string          `bson:"paymentMethod"`
	Status         string          `bson:"status"`
	CreatedAt      time.Time       `bson:"createdAt"`
}

type OrderItem struct {
	MenuItemID string `bson:"menuItemId"`
	Name       string `bson:"name"`
	Quantity   int    `bson:"quantity"`
	Price      int    `bson:"price"`
	FinalPrice int    `bson:"finalPrice"`
}

type Address struct {
	ID         string  `bson:"id"`
	Name       string  `bson:"name"`
	Phone      string  `bson:"phone"`
	Annotation string  `bson:"annotation"`
	Address    string  `bson:"address"`
	House      string  `bson:"house"`
	Area       string  `bson:"area,omitempty"`
	City       string  `bson:"city,omitempty"`
	Landmark   string  `bson:"landmark"`
	Lat        float64 `bson:"lat"`
	Lng        float64 `bson:"lng"`
}
