package badger

import (
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/threeaccents/cache"
)

type CacheStorage struct {
	DB *badger.DB
}

func (s *CacheStorage) Set(key string, value []byte) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (s *CacheStorage) SetTTL(key string, value []byte, ttl time.Duration) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), value).WithTTL(ttl)
		return txn.SetEntry(entry)
	})
}

func (s *CacheStorage) Get(key string) ([]byte, error) {
	var result []byte

	err := s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return cache.ErrKeyNotFound
			}
			return err
		}

		return item.Value(func(val []byte) error {
			result = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *CacheStorage) Delete(key string) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		if err := txn.Delete([]byte(key)); err != nil {
			if err == badger.ErrKeyNotFound {
				return cache.ErrKeyNotFound
			}
			return err
		}

		return nil
	})
}
