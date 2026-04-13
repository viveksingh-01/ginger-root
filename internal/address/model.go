package address

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Address struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	UserID     bson.ObjectID `bson:"userId"`
	Name       string        `bson:"name"`
	Phone      string        `bson:"phone"`
	Annotation string        `bson:"annotation"`
	Address    string        `bson:"address"`
	House      string        `bson:"house"`
	Area       string        `bson:"area,omitempty"`
	City       string        `bson:"city,omitempty"`
	Landmark   string        `bson:"landmark"`
	Lat        float64       `bson:"lat"`
	Lng        float64       `bson:"lng"`
	IsDefault  bool          `bson:"isDefault"`
}
