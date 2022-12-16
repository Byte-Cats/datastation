package pkg

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBConfig holds the configuration for a MongoDB connection.
type MongoDBConfig struct {
	URI      string
	Database string
	Username string
	Password string
}

// DefaultMongoDBConfig returns a default configuration for a MongoDB connection.
func DefaultMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI:      "mongodb://localhost:27017",
		Database: "test",
		Username: "",
		Password: "",
	}
}

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

// MongoDBOptions creates a *options.ClientOptions struct using a MongoDBConfig struct.
func MongoDBOptions(config MongoDBConfig) *options.ClientOptions {
	opts := options.Client().ApplyURI(config.URI)
	if config.Username != "" || config.Password != "" {
		opts.SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		})
	}
	return opts
}
