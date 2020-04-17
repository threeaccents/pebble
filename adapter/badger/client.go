package badger

import badger "github.com/dgraph-io/badger/v2"

func Open(dir string) (*badger.DB, error) {
	return badger.Open(badger.DefaultOptions(dir))
}
