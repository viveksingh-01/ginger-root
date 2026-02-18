package menu

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("menus")}
}

func (r *Repository) FindByRestaurantID(ctx context.Context, restaurantID string) (*Menu, error) {
	var menu Menu
	err := r.collection.FindOne(ctx, bson.M{"restaurantId": restaurantID}).Decode(&menu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrMenuNotFound
		}
		return nil, err
	}
	return &menu, nil
}
