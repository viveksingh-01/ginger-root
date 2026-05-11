package order

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
	return &Repository{collection: db.Collection("orders")}
}

func (r *Repository) Create(ctx context.Context, order *Order) error {
	_, err := r.collection.InsertOne(ctx, order)
	return err
}

// Returns the next 6-digit order id in [100000, 999999] (Uses an atomic counter in MongoDB)
// Unique per sequence until 900000 orders have been issued (then ids repeat).
func (r *Repository) NextOrderID(ctx context.Context) (int, error) {
	coll := r.collection.Database().Collection("counters")
	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var doc struct {
		Seq int64 `bson:"seq"`
	}
	err := coll.FindOneAndUpdate(ctx,
		bson.M{"_id": "order"},
		bson.M{"$inc": bson.M{"seq": 1}},
		opts,
	).Decode(&doc)
	if err != nil {
		return 0, err
	}

	const minID = 100000
	const span = 900000 // 100000..999999 inclusive
	if doc.Seq <= 0 {
		return minID, nil
	}
	return int((doc.Seq-1)%span) + minID, nil
}
