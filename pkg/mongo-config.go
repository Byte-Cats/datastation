package pkg

// MongoDBConfig holds the configuration for a MongoDB connection.
type MongoDBConfig struct {
	DatabaseConfig
	URI      string
	Database string
}

// DefaultMongoDBConfig returns a default configuration for a MongoDB connection.
func DefaultMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		DatabaseConfig: DefaultMongoDB(),
		URI:            "mongodb://localhost:27017",
		Database:       "test",
	}
}
