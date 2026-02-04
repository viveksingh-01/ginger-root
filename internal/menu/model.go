package menu

type OfferTags struct {
	Title    string `json:"title" bson:"title"`
	SubTitle string `json:"subTitle" bson:"subTitle"`
}

type Ratings struct {
	Rating        float64 `json:"rating" bson:"rating"`
	RatingCount   string  `json:"ratingCount" bson:"ratingCount"`
	RatingCountV2 int     `json:"ratingCountV2" bson:"ratingCountV2"`
}

type MenuItem struct {
	ID                 string    `json:"id" bson:"_id,omitempty"`
	Name               string    `json:"name" bson:"name"`
	Category           string    `json:"category" bson:"category"`
	Description        string    `json:"description" bson:"description"`
	ImageID            string    `json:"imageId" bson:"imageId"`
	InStock            int       `json:"inStock" bson:"inStock"`
	IsBestseller       bool      `json:"isBestseller" bson:"isBestseller"`
	IsVeg              bool      `json:"isVeg" bson:"isVeg"`
	Price              int       `json:"price" bson:"price"`
	FinalPrice         int       `json:"finalPrice" bson:"finalPrice"`
	ItemPriceStrikeOff bool      `json:"itemPriceStrikeOff" bson:"itemPriceStrikeOff"`
	OfferTags          OfferTags `json:"offerTags" bson:"offerTags"`
	PortionSize        string    `json:"portionSize" bson:"portionSize"`
	Ratings            Ratings   `json:"ratings" bson:"ratings"`
}

type Menu struct {
	RestaurantID string     `json:"restaurantId" bson:"restaurantId"`
	Items        []MenuItem `json:"items" bson:"items"`
}
