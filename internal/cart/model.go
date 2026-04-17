package cart

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CartItem struct {
	MenuItemID string `bson:"menuItemId"`
	Quantity   int    `bson:"quantity"`
}

type Cart struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	UserID       string        `bson:"userId,omitempty"`
	GuestID      string        `bson:"guestId,omitempty"`
	RestaurantID string        `bson:"restaurantId"`
	AddressID    string        `bson:"addressId"`
	Items        []CartItem    `bson:"cartItems"`
	CreatedAt    time.Time     `bson:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
}
