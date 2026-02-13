package restaurantmenu

import (
	"github.com/viveksingh-01/ginger-root/internal/menu"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
)

type RestaurantMenuResponse struct {
	Details *restaurant.Restaurant `json:"details"`
	Menu    []menu.MenuItem        `json:"menu"`
}
