package cart

type CartItem struct {
	MenuItemID string `bson:"menuItemId"`
	Quantity   int    `bson:"quantity"`
}

type Cart struct {
	ID           string     `bson:"_id,omitempty"`
	UserID       string     `bson:"userId,omitempty"`
	GuestID      string     `bson:"guestId,omitempty"`
	RestaurantID string     `bson:"restaurantId"`
	AddressID    string     `bson:"addressId"`
	Items        []CartItem `bson:"cartItems"`
}
