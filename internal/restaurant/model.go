package restaurant

import "go.mongodb.org/mongo-driver/v2/bson"

type SLA struct {
	DeliveryTime   int     `bson:"delivery_time" json:"deliveryTime"`
	LastMileTravel float64 `bson:"last_mile_travel" json:"lastMileTravel"`
	SLAString      string  `bson:"sla_string" json:"slaString"`
}

type Restaurant struct {
	ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name               string        `bson:"name" json:"name"`
	CloudinaryImageID  string        `bson:"cloudinary_image_id" json:"cloudinaryImageId"`
	Locality           string        `bson:"locality" json:"locality"`
	AreaName           string        `bson:"area_name" json:"areaName"`
	CostForTwo         string        `bson:"cost_for_two" json:"costForTwo"`
	Cuisines           []string      `bson:"cuisines" json:"cuisines"`
	AvgRating          float64       `bson:"avg_rating" json:"avgRating"`
	AvgRatingString    string        `bson:"avg_rating_string" json:"avgRatingString"`
	TotalRatingsString string        `bson:"total_ratings_string" json:"totalRatingsString"`
	Veg                bool          `bson:"veg" json:"veg"`
	SLA                SLA           `bson:"sla" json:"sla"`
}
