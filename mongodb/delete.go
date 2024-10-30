package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) DeleteAllUnavailable() (int64, error) {
	collection := db.client.Database(dbName).Collection(collectionName)

	filter := bson.D{{Key: "available", Value: false}}
	result, err := collection.DeleteMany(db.ctx, filter)

	return result.DeletedCount, err
}

func (db *DB) DeleteAllOverpriced() (int64, error) {
	collection := db.client.Database(dbName).Collection(collectionName)

	filter := bson.D{
		{
			Key: "$or", Value: bson.A{
				bson.D{{Key: "price", Value: bson.D{{Key: "$gt", Value: 200}}}},
				bson.D{{Key: "availability", Value: false}},
			},
		},
	}

	result, err := collection.DeleteMany(db.ctx, filter)

	return result.DeletedCount, err
}
