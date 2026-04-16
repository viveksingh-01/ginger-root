package cart

import (
	"context"

	"github.com/viveksingh-01/ginger-root/internal/auth"
	"github.com/viveksingh-01/ginger-root/internal/menu"
	"github.com/viveksingh-01/ginger-root/internal/restaurant"
)

type Service struct {
	repo              *Repository
	menuService       *menu.Service
	restaurantService *restaurant.Service
	userService       *auth.Service
}

func NewService(r *Repository, m *menu.Service, rs *restaurant.Service, us *auth.Service) *Service {
	return &Service{
		repo:              r,
		menuService:       m,
		restaurantService: rs,
		userService:       us,
	}
}

func (s *Service) AddToCart(
	ctx context.Context,
	userID string,
	restaurantID string,
	addressID string,
	items []CartItem,
) (*Response, error) {

	u, err := s.userService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	cart, err := s.repo.FindByUserID(ctx, userID)
	if err != nil && err != ErrCartNotFound {
		return nil, err
	}

	if err == ErrCartNotFound {
		cart = &Cart{
			UserID:       userID,
			RestaurantID: restaurantID,
			AddressID:    addressID,
			Items:        items,
		}
	} else {
		if cart.RestaurantID != restaurantID {
			cart.Items = items
			cart.RestaurantID = restaurantID
		} else {
			cart.Items = MergeItems(cart.Items, items)
		}
	}

	if err := s.repo.Upsert(ctx, cart); err != nil {
		return nil, err
	}

	restaurant, _ := s.restaurantService.GetRestaurant(ctx, restaurantID)
	menuData, _ := s.menuService.GetMenu(ctx, restaurantID)

	var responseItems []Item
	totalItems := 0
	subtotal := 0.0
	discount := 0.0

	for _, ci := range cart.Items {
		mi := FindMenuItem(menuData.Items, ci.MenuItemID)
		if mi == nil {
			continue
		}

		itemTotal := float64(mi.Price * ci.Quantity)
		finalPrice := float64(mi.FinalPrice * ci.Quantity)

		responseItems = append(responseItems, Item{
			MenuItemID:        ci.MenuItemID,
			Name:              mi.Name,
			Quantity:          ci.Quantity,
			Total:             int(itemTotal),
			FinalPrice:        int(finalPrice),
			IsVeg:             BoolToInt(mi.IsVeg),
			CloudinaryImageID: mi.ImageID,
		})

		totalItems += ci.Quantity
		subtotal += itemTotal
		discount += (itemTotal - finalPrice)
	}

	delivery := 4900.0
	gst := subtotal * 0.05
	finalAmount := subtotal + delivery + gst - discount

	return &Response{
		StatusCode:    0,
		StatusMessage: "CART_UPDATED_SUCCESSFULLY",
		Data: &CartResponse{
			CartMeta: CartMeta{
				CartID:     cart.ID,
				EmailID:    u.Email,
				PhoneNo:    u.Phone,
				CodEnabled: true,
				AddressID:  addressID,
				RestaurantDetails: RestaurantDetails{
					ID:                restaurant.ID.Hex(),
					Name:              restaurant.Name,
					CloudinaryImageID: restaurant.CloudinaryImageID,
					SLA: SLA{
						SLAString: restaurant.SLA.SLAString,
					},
				},
			},
			CartDetails: CartDetails{
				Items:           responseItems,
				TotalItemsCount: totalItems,
				BillDetails: BillDetails{
					Subtotal:       subtotal,
					DeliveryCharge: delivery,
					DiscountAmount: discount,
					GST:            gst,
					FinalAmount:    finalAmount,
				},
			},
		},
	}, nil
}
