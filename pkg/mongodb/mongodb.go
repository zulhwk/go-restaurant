package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClientMongo() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Define clientOptions.
	clientOptions := options.Client()
	clientOptions.ApplyURI(os.Getenv("MONGODB_URI"))

	// Initial client mongo.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		client.Disconnect(context.Background())
		return nil, err
	}
	return client, nil
}

func GetDatabaseMongo(dbName string) (*mongo.Database, error) {
	client, _ := CreateClientMongo()
	return client.Database(dbName), nil
}

func GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	col, _ := GetDatabaseMongo(dbName)
	return col.Collection(collectionName), nil
}
