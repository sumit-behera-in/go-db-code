package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) SwapOneAvailability(name string, price float64, available bool) error {
	collection := db.client.Database(dbName).Collection(collectionName)

	// Define the update operation to set "availability" to a new value
	update := bson.D{{"$set", bson.D{{"available", !available}}}}

	filter := bson.D{
		{"name", name},
		{"price", price},
		{"available", available},
	}
	result := collection.FindOneAndUpdate(db.ctx, filter, update)
	return result.Err()

}

func (db *DB) UpdateAllNames(previousName string, newName string) (int64, error) {
	collection := db.client.Database(dbName).Collection(collectionName)

	// Define the filter to match documents by "name"
	filter := bson.D{{"name", previousName}}

	// Define the update operation to set "availability" to a new value
	update := bson.D{{"$set", bson.D{{"name", newName}}}}

	result, err := collection.UpdateMany(db.ctx, filter, update)

	return result.ModifiedCount, err
}
