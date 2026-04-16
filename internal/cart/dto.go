package cart

type AddToCartRequest struct {
	Cart struct {
		RestaurantID string     `json:"restaurantId"`
		AddressID    string     `json:"addressId"`
		CartItems    []CartItem `json:"cartItems"`
	} `json:"cart"`
}

type CartResponse struct {
	CartMeta    CartMeta    `json:"cartMeta"`
	CartDetails CartDetails `json:"cartDetails"`
}

type CartMeta struct {
	CartID            string            `json:"cartId"`
	EmailID           string            `json:"emailId"`
	PhoneNo           string            `json:"phoneNo"`
	RestaurantDetails RestaurantDetails `json:"restaurantDetails"`
	CodEnabled        bool              `json:"codEnabled"`
	AddressID         string            `json:"addressId"`
}

type RestaurantDetails struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	CloudinaryImageID string `json:"cloudinaryImageId"`
	SLA               SLA    `json:"sla"`
}

type SLA struct {
	SLAString string `json:"slaString"`
}

type CartDetails struct {
	Items           []Item      `json:"items"`
	TotalItemsCount int         `json:"totalItemsCount"`
	BillDetails     BillDetails `json:"billDetails"`
}

type Item struct {
	MenuItemID        string `json:"menuItemId"`
	Name              string `json:"name"`
	Quantity          int    `json:"quantity"`
	Total             int    `json:"total"`
	FinalPrice        int    `json:"finalPrice"`
	StrikeOffEnabled  bool   `json:"strikeOffEnabled"`
	IsVeg             int    `json:"isVeg"`
	CloudinaryImageID string `json:"cloudinaryImageId"`
}

type BillDetails struct {
	Subtotal        float64 `json:"subtotal"`
	DeliveryCharge  float64 `json:"deliveryCharge"`
	DiscountAmount  float64 `json:"discountAmount"`
	CartTotal       float64 `json:"cartTotal"`
	GST             float64 `json:"GST"`
	PackingCharges  float64 `json:"packingCharges"`
	TaxesAndCharges float64 `json:"taxesAndCharges"`
	TotalAmount     float64 `json:"totalAmount"`
	FinalAmount     float64 `json:"finalAmount"`
}

type Response struct {
	StatusCode    int           `json:"statusCode"`
	StatusMessage string        `json:"statusMessage"`
	Data          *CartResponse `json:"data,omitempty"`
}
