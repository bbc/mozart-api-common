package caching

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisCache struct {
	pool *redis.Pool
}

func (c *RedisCache) getConn() redis.Conn {
	if c.pool == nil {
		c.pool = &redis.Pool{
			MaxIdle: 30,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST"))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		}
	}

	return c.pool.Get()
}

func (c *RedisCache) Get(key string) (string, error) {
	conn := c.getConn()
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", fmt.Errorf("can't get redis key %q: %s", key, err)
	}

	return val, nil
}

func (c *RedisCache) Set(key string, data string, expiration time.Duration) (string, error) {
	conn := c.getConn()
	defer conn.Close()

	val, err := redis.String(conn.Do("SET", key, data, expiration))
	if err != nil {
		return "", fmt.Errorf("can't set redis key %q: %s", key, err)
	}

	return val, nil

}
