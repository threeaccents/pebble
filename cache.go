package cache

import "time"

type CacheStorage interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	SetTTL(key string, value []byte, ttl time.Duration) error
}
