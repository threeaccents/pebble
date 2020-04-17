package main

import (
	"log"

	"github.com/threeaccents/cache/adapter/badger"

	"github.com/threeaccents/cache/transport/grpc"
)

func main() {
	db, err := badger.Open("./db")
	if err != nil {
		panic(err)
	}

	storage := &badger.CacheStorage{
		DB: db,
	}

	log.Fatal(grpc.Serve(storage))
}
