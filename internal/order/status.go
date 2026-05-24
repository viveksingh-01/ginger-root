package order

import "time"

const (
	StatusPlaced    = "PLACED"
	StatusConfirmed = "CONFIRMED"
	StatusPreparing = "PREPARING"
	StatusReady     = "READY"
	StatusPickedUp  = "PICKED_UP"
	StatusOnTheWay  = "ON_THE_WAY"
	StatusDelivered = "DELIVERED"
)

type StatusStep struct {
	Status   string
	Title    string
	Subtitle string
	Delay    time.Duration
}

var statusPipeline = []StatusStep{
	{
		Status:   StatusPlaced,
		Title:    "Order placed",
		Subtitle: "We have received your order",
		Delay:    4 * time.Second,
	},
	{
		Status:   StatusConfirmed,
		Title:    "Restaurant confirmed",
		Subtitle: "The restaurant has accepted your order",
		Delay:    6 * time.Second,
	},
	{
		Status:   StatusPreparing,
		Title:    "Preparing your food",
		Subtitle: "Your order is being prepared",
		Delay:    8 * time.Second,
	},
	{
		Status:   StatusReady,
		Title:    "Ready for pickup",
		Subtitle: "Waiting for delivery partner",
		Delay:    6 * time.Second,
	},
	{
		Status:   StatusPickedUp,
		Title:    "Picked up",
		Subtitle: "Delivery partner has collected your order",
		Delay:    5 * time.Second,
	},
	{
		Status:   StatusOnTheWay,
		Title:    "On the way",
		Subtitle: "Your order is heading to you",
		Delay:    10 * time.Second,
	},
	{
		Status:   StatusDelivered,
		Title:    "Delivered",
		Subtitle: "Enjoy your meal!",
		Delay:    0,
	},
}
