package badger

import (
	"time"

	"github.com/oriiolabs/pebble"
	"github.com/rs/zerolog"

	"github.com/dgraph-io/badger/v2"
)

type CacheStorage struct {
	DB *badger.DB

	Log zerolog.Logger
}

func (s *CacheStorage) Set(key string, value []byte) error {
	s.Log.Info().Str("key", key).Msg("setting key")
	return s.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (s *CacheStorage) SetTTL(key string, value []byte, ttl time.Duration) error {
	s.Log.Info().Str("key", key).Msg("setting key")
	return s.DB.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), value).WithTTL(ttl)
		return txn.SetEntry(entry)
	})
}

func (s *CacheStorage) Get(key string) ([]byte, error) {
	var result []byte
	s.Log.Info().Str("key", key).Msg("getting key")

	err := s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return pebble.ErrKeyNotFound
			}
			return err
		}

		return item.Value(func(val []byte) error {
			result = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		s.Log.Info().Str("key", key).Msg("key not found")
		return nil, err
	}

	s.Log.Info().Str("key", key).Msg("key found")

	return result, nil
}

func (s *CacheStorage) Delete(key string) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		if err := txn.Delete([]byte(key)); err != nil {
			if err == badger.ErrKeyNotFound {
				return pebble.ErrKeyNotFound
			}
			return err
		}

		return nil
	})
}
