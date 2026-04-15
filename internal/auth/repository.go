package auth

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("users")}
}

// Create a User
func (r *Repository) Create(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// Retrieve User by Phone number
func (r *Repository) FindByPhone(ctx context.Context, phone string) (*User, error) {
	var user *User
	err := r.collection.FindOne(ctx, bson.M{"phone": phone}).Decode((&user))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) FindByID(ctx context.Context, id bson.ObjectID) (*User, error) {
	var user User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
