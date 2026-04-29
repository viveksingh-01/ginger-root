package order

import "github.com/viveksingh-01/ginger-root/internal/address"

func ToOrderResponse(o *Order) *OrderResponse {
	return &OrderResponse{
		ID:             o.ID,
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
		Items: o.Items,
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
