package order

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("orders")}
}

func (r *Repository) Create(ctx context.Context, order *Order) error {
	_, err := r.collection.InsertOne(ctx, order)
	return err
}
