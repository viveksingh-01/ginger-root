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

	// Calculate totals
	var items []OrderItem
	subtotal := 0.0
	discount := 0.0

	for _, ci := range cartData.Items {
		mi := cart.FindMenuItem(menuData.Items, ci.MenuItemID)
		if mi == nil {
			continue
		}

		itemTotal := float64(mi.Price * ci.Quantity)
		finalPrice := float64(mi.FinalPrice * ci.Quantity)

		items = append(items, OrderItem{
			MenuItemID: ci.MenuItemID,
			Name:       mi.Name,
			Quantity:   ci.Quantity,
			Price:      mi.Price,
			FinalPrice: mi.FinalPrice,
		})

		subtotal += itemTotal
		discount += (itemTotal - finalPrice)
	}

	delivery := 4900.0
	gst := subtotal * 0.05
	finalAmount := subtotal + delivery + gst - discount

	orderID, err := s.repo.NextOrderID(ctx)
	if err != nil {
		return nil, err
	}

	// Create order
	order := &Order{
		ID:             uuid.NewString(),
		OrderID:        orderID,
		UserID:         userID,
		RestaurantID:   r.ID.Hex(),
		RestaurantName: r.Name,
		Address:        *address,
		Items:          items,
		Subtotal:       subtotal,
		Delivery:       delivery,
		GST:            gst,
		Discount:       discount,
		FinalAmount:    finalAmount,
		PaymentMethod:  paymentMethod,
		Status:         "PLACED",
		CreatedAt:      time.Now(),
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return nil, err
	}

	// Delete the cart once order is created
	if err := s.cartRepo.DeleteByID(ctx, cartID); err != nil {
		return nil, err
	}

	return order, nil
}

// Each stage is persisted and sent over SSE; delays between stages simulate live progression until DELIVERED.
func (s *Service) StreamOrderStatus(
	ctx context.Context,
	userID string,
	orderID int,
	emit func(OrderStatusEvent) error,
) error {
	order, err := s.repo.FindByOrderID(ctx, orderID)
	if err != nil {
		return err
	}
	if order.UserID != userID {
		return ErrOrderForbidden
	}

	total := len(statusPipeline)
	start := StatusIndex(order.Status)

	for i := start; i < total; i++ {
		step := statusPipeline[i]
		ev := OrderStatusEvent{
			OrderID:        orderID,
			RestaurantName: order.RestaurantName,
			Status:         step.Status,
			Title:          step.Title,
			Subtitle:       step.Subtitle,
			Step:           i + 1,
			TotalSteps:     total,
			IsTerminal:     step.Status == StatusDelivered,
			UpdatedAt:      time.Now().UTC().Format(time.RFC3339),
		}

		if err := emit(ev); err != nil {
			return err
		}

		if step.Status != order.Status {
			if err := s.repo.UpdateStatus(ctx, orderID, step.Status); err != nil {
				return err
			}
			order.Status = step.Status
		}

		if step.Status == StatusDelivered {
			return nil
		}

		if i < total-1 && step.Delay > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(step.Delay):
			}
		}
	}

	return nil
}
