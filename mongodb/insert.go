package mongodb

import (
	"fmt"

	"github.com/sumit-behera-in/go-db-code/structs"
)

func (db *DB) Insert(product structs.Product) (string, error) {
	// collection == table
	collection := db.client.Database(dbName).Collection(collectionName)
	result, err := collection.InsertOne(db.ctx, product)
	return fmt.Sprint(result.InsertedID), err
}
