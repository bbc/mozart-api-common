package caching

import "time"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, data string, expiration time.Duration) (string, error)
}
