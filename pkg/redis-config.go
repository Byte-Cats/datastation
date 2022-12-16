package pkg

import "github.com/go-redis/redis"

type RedisConfig struct {
	DatabaseConfig
	Address  string
	DB       int
	PoolSize int
}

// RedisOptions creates a *redis.Options struct using a RedisConfig struct.
func RedisOptions(config RedisConfig) *redis.Options {
	return &redis.Options{
		Addr:     config.Address,
		DB:       config.DB,
		Password: config.Password,
		PoolSize: config.PoolSize,
	}
}

// DefaultRedisConfig returns a default configuration for a Redis connection.
func DefaultRedisConfig() RedisConfig {
	redConf := DefaultRedis()
	return RedisConfig{
		DatabaseConfig: redConf,
		Address:        "localhost:6379",
		DB:             0,
		PoolSize:       10,
	}
}
