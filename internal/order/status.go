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
	Status     string
	Title      string
	Subtitle   string
	Delay      time.Duration
	ETAMinutes int // estimated minutes until delivery at this stage
}

var statusPipeline = []StatusStep{
	{
		Status:     StatusPlaced,
		Title:      "Order placed",
		Subtitle:   "We have received your order",
		Delay:      2 * time.Second,
		ETAMinutes: 38,
	},
	{
		Status:     StatusConfirmed,
		Title:      "Restaurant confirmed",
		Subtitle:   "The restaurant has accepted your order",
		Delay:      4 * time.Second,
		ETAMinutes: 36,
	},
	{
		Status:     StatusPreparing,
		Title:      "Preparing your food",
		Subtitle:   "Your order is being prepared",
		Delay:      12 * time.Second,
		ETAMinutes: 32,
	},
	{
		Status:     StatusReady,
		Title:      "Ready for pickup",
		Subtitle:   "Waiting for delivery partner",
		Delay:      4 * time.Second,
		ETAMinutes: 20,
	},
	{
		Status:     StatusPickedUp,
		Title:      "Picked up",
		Subtitle:   "Delivery partner has collected your order",
		Delay:      2 * time.Second,
		ETAMinutes: 16,
	},
	{
		Status:     StatusOnTheWay,
		Title:      "On the way",
		Subtitle:   "Your order is heading to you",
		Delay:      14 * time.Second,
		ETAMinutes: 14,
	},
	{
		Status:     StatusDelivered,
		Title:      "Delivered",
		Subtitle:   "Enjoy your meal!",
		Delay:      0,
		ETAMinutes: 0,
	},
}

func StatusIndex(status string) int {
	for i, step := range statusPipeline {
		if step.Status == status {
			return i
		}
	}
	return 0
}
