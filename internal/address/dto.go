package address

type CreateAddressRequest struct {
	Name       string  `json:"name" binding:"required"`
	Phone      string  `json:"phone" binding:"required,len=10,numeric"`
	Annotation string  `json:"annotation" binding:"required"`
	Address    string  `json:"address" binding:"required"`
	House      string  `json:"house" binding:"required"`
	Area       string  `json:"area" binding:"required"`
	City       string  `json:"city" binding:"required"`
	Landmark   string  `json:"landmark,omitempty"`
	Lat        float64 `json:"lat" binding:"required"`
	Lng        float64 `json:"lng" binding:"required"`
}

type AddressResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Phone      string  `json:"phone"`
	Annotation string  `json:"annotation"`
	Address    string  `json:"address"`
	House      string  `json:"house"`
	Area       string  `json:"area"`
	City       string  `json:"city"`
	Landmark   string  `json:"landmark,omitempty"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
}

type Response struct {
	StatusCode    int               `json:"statusCode"`
	StatusMessage string            `json:"statusMessage"`
	Data          []AddressResponse `json:"data,omitempty"`
}
