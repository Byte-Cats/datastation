package pkg

import "github.com/go-redis/redis"

// NewRedisClient creates a new Redis client using the provided configuration.
func NewRedisClient(config RedisConfig) *redis.Client {
	return redis.NewClient(RedisOptions(config))
}

// RedisPing pings the Redis server to test the connection.
func RedisPing(client *redis.Client) error {
	return client.Ping().Err()
}

// RedisClose closes the connection to the Redis server.
func RedisClose(client *redis.Client) error {
	return client.Close()
}

// Set sets the value of a key in Redis.
func Set(client *redis.Client, key string, value interface{}) error {
	return client.Set(key, value, 0).Err()
}

// RedisGet gets the value of a key in Redis.
func RedisGet(client *redis.Client, key string) (interface{}, error) {
	return client.Get(key).Result()
}

// RedisDelete deletes a key in Redis.
func RedisDelete(client *redis.Client, key string) error {
	return client.Del(key).Err()
}

// RedisExists checks if a key exists in Redis.
func RedisExists(client *redis.Client, key string) (int64, error) {
	return client.Exists(key).Result()
}

// RedisHSet sets the value of a field in a hash in Redis.
func RedisHSet(client *redis.Client, key string, field string, value interface{}) error {
	return client.HSet(key, field, value).Err()
}

// RedisHGet gets the value of a field in a hash in Redis.
func RedisHGet(client *redis.Client, key string, field string) (interface{}, error) {
	return client.HGet(key, field).Result()
}

// RedisHExists checks if a field exists in a hash in Redis.
func RedisHExists(client *redis.Client, key string, field string) (bool, error) {
	return client.HExists(key, field).Result()
}

// RedisSAdd adds a member to a set in Redis.
func RedisSAdd(client *redis.Client, key string, member interface{}) error {
	return client.SAdd(key, member).Err()
}
