package restaurantmenu

import (
	"context"

	"github.com/viveksingh-01/ginger-root/internal/menu"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
)

type Service struct {
	restaurantService *restaurant.Service
	menuService       *menu.Service
}

func NewService(rs *restaurant.Service, ms *menu.Service) *Service {
	return &Service{
		restaurantService: rs,
		menuService:       ms,
	}
}

func (s *Service) GetRestaurantMenu(ctx context.Context, restaurantId string) (*RestaurantMenuResponse, error) {
	restaurant, err := s.restaurantService.GetRestaurant(ctx, restaurantId)
	if err != nil {
		return nil, err
	}
	menu, err := s.menuService.GetMenu(ctx, restaurantId)
	if err != nil {
		return nil, err
	}
	return &RestaurantMenuResponse{
		Details: restaurant,
		Menu:    menu.Items,
	}, nil
}
