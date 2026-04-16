package cart

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("cart")}
}

func (r *Repository) FindByUserID(ctx context.Context, userID string) (*Cart, error) {
	var cart Cart
	err := r.collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrCartNotFound
		}
		return nil, err
	}
	return &cart, nil
}

func (r *Repository) Upsert(ctx context.Context, cart *Cart) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"userId": cart.UserID},
		bson.M{"$set": cart},
	)
	return err
}
