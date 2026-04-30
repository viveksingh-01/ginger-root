package order

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/viveksingh-01/ginger-root/internal/address"
	"github.com/viveksingh-01/ginger-root/internal/cart"
	"github.com/viveksingh-01/ginger-root/internal/menu"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
)

type Service struct {
	repo              *Repository
	cartRepo          *cart.Repository
	menuService       *menu.Service
	restaurantService *restaurant.Service
	addressService    *address.Service
}

func NewService(
	r *Repository,
	cartRepo *cart.Repository,
	m *menu.Service,
	rs *restaurant.Service,
	as *address.Service,
) *Service {
	return &Service{
		repo:              r,
		cartRepo:          cartRepo,
		menuService:       m,
		restaurantService: rs,
		addressService:    as,
	}
}

func (s *Service) PlaceOrder(ctx context.Context, userID, cartID, addressID, paymentMethod string) (*Order, error) {
	// Fetch cart
	cartData, err := s.cartRepo.FindCartByID(ctx, cartID)
	if err != nil {
		return nil, err
	}

	// Fetch restaurant
	r, err := s.restaurantService.GetRestaurant(ctx, cartData.RestaurantID)
	if err != nil {
		return nil, err
	}

	// Fetch menu
	menuData, err := s.menuService.GetMenu(ctx, cartData.RestaurantID)
	if err != nil || menuData == nil {
		menuData = &menu.Menu{Items: []menu.MenuItem{}}
	}

	// Fetch address
	address, err := s.addressService.GetAddress(ctx, addressID)
	if err != nil {
		return nil, err
	}
	if userID != "" && address.UserID.Hex() != userID {
		return nil, errors.New("address does not belong to user")
	}

	// Create order
	order := &Order{
		ID:             uuid.NewString(),
		UserID:         userID,
		RestaurantID:   r.ID.Hex(),
		RestaurantName: r.Name,
		PaymentMethod:  paymentMethod,
		Address:        *address,
		Status:         "PLACED",
		CreatedAt:      time.Now(),
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
