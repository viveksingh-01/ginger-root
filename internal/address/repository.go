package address

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("addresses")}
}

func (r *Repository) Create(ctx context.Context, address *Address) (*Address, error) {
	res, err := r.collection.InsertOne(ctx, address)
	if err != nil {
		return nil, err
	}

	address.ID = res.InsertedID.(bson.ObjectID)
	return address, nil
}

func (r *Repository) GetByUserID(ctx context.Context, userId bson.ObjectID) ([]*Address, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"userId": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var addresses []*Address
	for cursor.Next(ctx) {
		var addr Address
		if err := cursor.Decode(&addr); err != nil {
			return nil, err
		}
		addresses = append(addresses, &addr)
	}

	return addresses, nil
}

var ErrAddressNotFound = errors.New("address not found")

func (r *Repository) FindByID(ctx context.Context, addressID string) (*Address, error) {
	var addr Address

	err := r.collection.FindOne(ctx, bson.M{"id": addressID}).Decode(&addr)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrAddressNotFound
		}
		return nil, err
	}

	return &addr, nil
}
