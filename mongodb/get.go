package mongodb

import (
	"github.com/sumit-behera-in/go-db-code/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) GetAll() ([]structs.Product, error) {
	collection := db.client.Database(dbName).Collection(collectionName)

	// 	pass mongodb options to find function
	findOption := options.Find()

	var dataCollection []structs.Product
	rows, err := collection.Find(db.ctx, bson.M{}, findOption)
	if err != nil {
		return dataCollection, err
	}

	defer rows.Close(db.ctx)

	var data structs.Product

	for rows.Next(db.ctx) {
		err = rows.Decode(&data)
		if err != nil {
			return dataCollection, err
		}
		dataCollection = append(dataCollection, data)
	}

	return dataCollection, nil

}
