package restaurant

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (r *Repository) List(ctx context.Context, skip, limit int64, f Filter, sort Sort) ([]Restaurant, error) {
	// Creates MongoDB query options. (Think of it like: “Extra instructions for MongoDB”)
	opts := options.Find()
	// Limits how many documents MongoDB returns
	opts.SetLimit(limit)
	// Skips offset number of documents
	opts.SetSkip(skip)
	// Set sorting
	opts.SetSort(bson.D{{Key: sort.Field, Value: sort.Order}})

	// bson.M{} is an empty filter, meaning: “Return all documents in the collection”
	// How this works?
	// Request: GET /restaurants?veg=true -->
	// Code: filter := bson.M{ "veg": true } -->
	// Mongo: db.restaurants.find({ veg: true })
	filter := bson.M{}
	if f.Veg != nil {
		filter["veg"] = *f.Veg
	}

	// Average-rating filter
	if f.MinRating != nil {
		// Apply 'greater than or equal to' condition for min-rating in Mongo query
		filter["avg_rating"] = bson.M{
			"$gte": *f.MinRating,
		}
	}

	// Find - queries the MongoDB collection with context, filter and options
	// The result is a cursor, which allows iterating over the documents
	cursor, err := r.collection.Find(ctx, filter, opts)
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

func (r *Repository) FindByID(ctx context.Context, restaurantId string) (*Restaurant, error) {
	var restaurant Restaurant
	objID, err := bson.ObjectIDFromHex(restaurantId)
	if err != nil {
		return nil, err // invalid ObjectID string
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&restaurant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrRestaurantNotFound
		}
		return nil, err
	}
	return &restaurant, nil
}

func (r *Repository) Search(ctx context.Context, query string) ([]Restaurant, error) {
	filter := bson.M{
		"name": bson.M{"$regex": query},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var restaurants []Restaurant
	if err = cursor.All(ctx, &restaurants); err != nil {
		return nil, err
	}
	return restaurants, nil
}
