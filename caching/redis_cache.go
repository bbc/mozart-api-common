package caching

import (
	"os"
	"time"

	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/gopkg.in/redis.v3"
)

type RedisCache struct {
	client *redis.Client
}

func (c *RedisCache) getClient() *redis.Client {
	if c.client == nil {
		c.client = redis.NewClient(&redis.Options{
			Addr:        os.Getenv("REDIS_HOST"),
			Password:    "", // no password set
			DB:          0,  // use default DB
			DialTimeout: 3 * time.Second,
			PoolSize:    300,
		})
	}

	return c.client
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.getClient().Get(key).Result()
}

func (c *RedisCache) Set(key string, data string, expiration time.Duration) (string, error) {
	return c.getClient().Set(key, data, expiration).Result()
}
