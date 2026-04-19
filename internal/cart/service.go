package cart

import (
	"context"
	"encoding/hex"

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
	userID, guestID, restaurantID, addressID string,
	items []CartItem,
) (*Response, error) {

	var u *auth.User
	var err error

	if userID != "" {
		u, err = s.userService.GetUser(ctx, userID)
		if err != nil {
			return nil, err
		}
	}

	cart, err := s.repo.FindCart(ctx, userID, guestID)
	if err != nil && err != ErrCartNotFound {
		return nil, err
	}

	if err == ErrCartNotFound {
		cart = &Cart{
			UserID:       userID,
			GuestID:      guestID,
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

	r, err := s.restaurantService.GetRestaurant(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	menuData, err := s.menuService.GetMenu(ctx, restaurantID)
	if err != nil || menuData == nil {
		menuData = &menu.Menu{RestaurantID: restaurantID, Items: []menu.MenuItem{}}
	}

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

	email, phone := "", ""
	if u != nil {
		email = u.Email
		phone = u.Phone
	}

	return &Response{
		StatusCode:    0,
		StatusMessage: "CART_UPDATED_SUCCESSFULLY",
		Data: &CartResponse{
			CartMeta: CartMeta{
				CartID:     cart.ID.Hex(),
				EmailID:    email,
				PhoneNo:    phone,
				CodEnabled: true,
				AddressID:  addressID,
				RestaurantDetails: RestaurantDetails{
					ID:                r.ID.Hex(),
					Name:              r.Name,
					CloudinaryImageID: r.CloudinaryImageID,
					SLA: SLA{
						SLAString: r.SLA.SLAString,
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

func (s *Service) GetCart(
	ctx context.Context,
	userID, guestID string,
) (*Response, error) {

	cart, err := s.repo.FindCart(ctx, userID, guestID)
	if err != nil {
		if err == ErrCartNotFound {
			return s.emptyCartResponse(), nil
		}
		return nil, err
	}

	return s.buildCartResponse(ctx, cart)
}

func (s *Service) emptyCartResponse() *Response {
	return &Response{
		StatusCode:    0,
		StatusMessage: "SUCCESS",
		Data: &CartResponse{
			CartMeta: CartMeta{},
			CartDetails: CartDetails{
				Items:           []Item{},
				TotalItemsCount: 0,
				BillDetails:     BillDetails{},
			},
		},
	}
}

func (s *Service) buildCartResponse(
	ctx context.Context,
	cart *Cart,
) (*Response, error) {
	if !isValidObjectIDHex(cart.RestaurantID) {
		return s.emptyCartResponse(), nil
	}

	var (
		u     *auth.User
		err   error
		email string
		phone string
	)

	if cart.UserID != "" {
		u, err = s.userService.GetUser(ctx, cart.UserID)
		if err != nil {
			return nil, err
		}
		email = u.Email
		phone = u.Phone
	}

	r, err := s.restaurantService.GetRestaurant(ctx, cart.RestaurantID)
	if err != nil {
		if err == restaurant.ErrRestaurantNotFound {
			return s.emptyCartResponse(), nil
		}
		return nil, err
	}

	menuData, err := s.menuService.GetMenu(ctx, cart.RestaurantID)
	if err != nil || menuData == nil {
		menuData = &menu.Menu{RestaurantID: cart.RestaurantID, Items: []menu.MenuItem{}}
	}

	responseItems := []Item{}
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
		StatusMessage: "SUCCESS",
		Data: &CartResponse{
			CartMeta: CartMeta{
				CartID:     cart.ID.Hex(),
				EmailID:    email,
				PhoneNo:    phone,
				CodEnabled: true,
				AddressID:  cart.AddressID,
				RestaurantDetails: RestaurantDetails{
					ID:                r.ID.Hex(),
					Name:              r.Name,
					CloudinaryImageID: r.CloudinaryImageID,
					SLA: SLA{
						SLAString: r.SLA.SLAString,
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

func isValidObjectIDHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}
