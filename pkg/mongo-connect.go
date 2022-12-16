package pkg

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// ConnectMongoDB establishes a connection to a MongoDB server.
func ConnectMongoDB(config MongoDBConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, MongoDBOptions(config))
	if err != nil {
		return nil, err
	}

	err = MongoDBPing(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// DisconnectMongoDB closes the connection to a MongoDB server.
func DisconnectMongoDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.Disconnect(ctx)
}

// MongoDBPing pings the MongoDB server to test the connection.
func MongoDBPing(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return client.Ping(ctx, nil)
}
