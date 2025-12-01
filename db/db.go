package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(dbUri, dbName string) (*mongo.Client, *mongo.Database, error) {
	//
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbUri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	// Ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	db := client.Database(dbName)

	return client, db, nil
}
