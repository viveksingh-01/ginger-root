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

func (r *Repository) FindCart(ctx context.Context, userID, guestID string) (*Cart, error) {
	filter := bson.M{}

	if userID != "" {
		filter["userId"] = userID
	} else {
		filter["guestId"] = guestID
	}

	var cart Cart
	err := r.collection.FindOne(ctx, filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrCartNotFound
		}
		return nil, err
	}
	return &cart, nil
}

func (r *Repository) Upsert(ctx context.Context, cart *Cart) error {
	filter := bson.M{}
	if cart.UserID != "" {
		filter["userId"] = cart.UserID
	} else {
		filter["guestId"] = cart.GuestID
	}

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": cart},
	)
	return err
}
