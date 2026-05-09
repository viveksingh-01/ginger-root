package cart

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("carts")}
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

	update := bson.M{
		"$set": bson.M{
			"userId":       cart.UserID,
			"guestId":      cart.GuestID,
			"restaurantId": cart.RestaurantID,
			"cartItems":    cart.Items,
			"updatedAt":    cart.UpdatedAt,
		},
		"$setOnInsert": bson.M{
			"createdAt": cart.CreatedAt,
		},
	}

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		update,
		options.UpdateOne().SetUpsert(true),
	)
	return err
}

func (r *Repository) FindCartByID(ctx context.Context, cartID string) (*Cart, error) {
	objID, err := bson.ObjectIDFromHex(cartID)
	if err != nil {
		return nil, err
	}

	var cart Cart
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrCartNotFound
		}
		return nil, err
	}
	return &cart, nil
}

func (r *Repository) DeleteByID(ctx context.Context, cartID string) error {
	objID, err := bson.ObjectIDFromHex(cartID)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
