package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) SwapOneAvailability(name string, price float64, available bool) error {
	collection := db.client.Database(dbName).Collection(collectionName)

	// Define the update operation to set "availability" to a new value
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "available", Value: !available}}}}

	filter := bson.D{
		{Key: "name", Value: name},
		{Key: "price", Value: price},
		{Key: "available", Value: available},
	}
	result := collection.FindOneAndUpdate(db.ctx, filter, update)
	return result.Err()

}

func (db *DB) UpdateAllNames(previousName string, newName string) (int64, error) {
	collection := db.client.Database(dbName).Collection(collectionName)

	// Define the filter to match documents by "name"
	filter := bson.D{{Key: "name", Value: previousName}}

	// Define the update operation to set "availability" to a new value
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newName}}}}

	result, err := collection.UpdateMany(db.ctx, filter, update)

	return result.ModifiedCount, err
}
