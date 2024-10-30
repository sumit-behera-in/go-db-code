package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func New() (DB, error) {

	db := DB{}
	var err error

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s%s", protocol, connetionURL))
	db.ctx, db.cancel = context.WithTimeout(context.Background(), 10*time.Second)
	db.client, err = mongo.Connect(db.ctx, clientOptions)
	if err != nil {
		return db, err
	}

	err = db.client.Ping(db.ctx, nil)
	if err != nil {
		return db, err
	}

	fmt.Println("Connected to MongoDB!")

	return db, nil
}
