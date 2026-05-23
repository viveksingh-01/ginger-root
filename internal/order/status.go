package order

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
	Status string
	Title  string
}

var statusPipeline = []StatusStep{
	{
		Status: StatusPlaced,
		Title:  "Order placed",
	},
	{
		Status: StatusConfirmed,
		Title:  "Restaurant confirmed",
	},
	{
		Status: StatusPreparing,
		Title:  "Preparing your food",
	},
	{
		Status: StatusReady,
		Title:  "Ready for pickup",
	},
	{
		Status: StatusPickedUp,
		Title:  "Picked up",
	},
	{
		Status: StatusOnTheWay,
		Title:  "On the way",
	},
	{
		Status: StatusDelivered,
		Title:  "Delivered",
	},
}
