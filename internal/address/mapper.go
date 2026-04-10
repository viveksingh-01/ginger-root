package address

func ToAddressResponse(a *Address) *AddressResponse {
	return &AddressResponse{
		ID:         a.ID.Hex(),
		Name:       a.Name,
		Phone:      a.Phone,
		Annotation: a.Annotation,
		Address:    a.Address,
		House:      a.House,
		Area:       a.Area,
		City:       a.City,
		Landmark:   a.Landmark,
		Lat:        a.Lat,
		Lng:        a.Lng,
	}
}
