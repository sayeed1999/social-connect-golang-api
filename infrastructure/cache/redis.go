package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient interface {
	Init()
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type cacheClient struct {
	client            *redis.Client
	defaultExpiration time.Duration
}

func NewCacheService() CacheClient {
	return &cacheClient{
		defaultExpiration: 10 * time.Minute,
	}
}

func (c *cacheClient) Init() {

	if c.client == nil {
		c.client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // No password set
			DB:       0,  // Use default DB
			Protocol: 2,  // Connection protocol
		})
	}
}

// Set stores a key-value pair in Redis with an expiration time
func (c *cacheClient) Set(ctx context.Context, key string, value string) error {
	return c.client.Set(ctx, key, value, c.defaultExpiration).Err()
}

// Get retrieves a value from Redis by key
func (c *cacheClient) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Delete removes a key-value pair from Redis
func (c *cacheClient) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Exists checks if a key exists in Redis
func (c *cacheClient) Exists(ctx context.Context, key string) (bool, error) {
	result, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return result > 0, nil
}
