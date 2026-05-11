package order

import (
	"github.com/viveksingh-01/ginger-root/internal/address"
	"github.com/viveksingh-01/ginger-root/internal/cart"
)

func ToOrderResponse(o *Order) *OrderResponse {
	return &OrderResponse{
		ID:             o.ID,
		OrderID:        o.OrderID,
		RestaurantID:   o.RestaurantID,
		RestaurantName: o.RestaurantName,
		Address: address.AddressResponse{
			ID:         o.Address.ID.Hex(),
			Name:       o.Address.Name,
			Phone:      o.Address.Phone,
			Annotation: o.Address.Annotation,
			Address:    o.Address.Address,
			House:      o.Address.House,
			Landmark:   o.Address.Landmark,
			Lat:        o.Address.Lat,
			Lng:        o.Address.Lng,
		},
		Items: toCartItems(o.Items),
		BillDetails: BillDetails{
			Subtotal:       o.Subtotal,
			DeliveryCharge: o.Delivery,
			GST:            o.GST,
			Discount:       o.Discount,
			FinalAmount:    o.FinalAmount,
		},
		Status: o.Status,
	}
}

func toCartItems(items []OrderItem) []cart.Item {
	out := make([]cart.Item, 0, len(items))
	for _, it := range items {
		out = append(out, cart.Item{
			MenuItemID: it.MenuItemID,
			Name:       it.Name,
			Quantity:   it.Quantity,
			Total:      it.Price * it.Quantity,
			FinalPrice: it.FinalPrice,
		})
	}
	return out
}
