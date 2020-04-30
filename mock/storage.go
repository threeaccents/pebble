package mock

import (
	"time"
)

type CacheStorage struct {
}

func (s *CacheStorage) Set(key string, value []byte) error {
	return nil
}

func (s *CacheStorage) SetTTL(key string, value []byte, ttl time.Duration) error {
	return nil
}

func (s *CacheStorage) Get(key string) ([]byte, error) {
	var result []byte

	return result, nil
}

func (s *CacheStorage) Delete(key string) error {
	return nil
}
