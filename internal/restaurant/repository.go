package restaurant

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Purpose:
// This struct implements the Repository pattern, which:
// Encapsulates database access logic
// Keeps MongoDB details out of your business logic
// Think of it as: “The only object allowed to touch the database.”
// Field: collection - a pointer to a MongoDB collection (restaurants)
type Repository struct {
	collection *mongo.Collection
}

// Implements Dependency Injection, a key companion to the Repository pattern.
// Makes the repository easy to initialize
// Useful for testing or switching databases
func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("restaurants")}
}

func (r *Repository) List(ctx context.Context) ([]Restaurant, error) {
	// Find queries the MongoDB collection
	// bson.M{} is an empty filter, meaning: “Return all documents in the collection”
	// The result is a cursor, which allows iterating over the documents
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	// Ensures the cursor is closed when the function finishes
	// Prevents memory leaks and open database resources
	defer cursor.Close(ctx)

	// Reads all documents from the cursor
	// Converts them into Restaurant structs
	// Stores them in the restaurants slice
	var restaurants []Restaurant
	if err := cursor.All(ctx, &restaurants); err != nil {
		return nil, err
	}

	return restaurants, nil
}
