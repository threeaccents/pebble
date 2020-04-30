package main

import (
	"flag"
	"log"

	"github.com/oriiolabs/pebble/adapter/badger"

	"github.com/oriiolabs/pebble/transport/grpc"
)

var (
	dbDirPtr = flag.String("db-dir", "./db", "the directory were the cache is stored")
	portPtr  = flag.String("port", ":4200", "the port were to run the grpc server")
)

func main() {
	flag.Parse()

	db, err := badger.Open(*dbDirPtr)
	if err != nil {
		panic(err)
	}

	storage := &badger.CacheStorage{
		DB: db,
	}

	s := grpc.NewServer(
		storage,
		grpc.ServerPort(*portPtr),
	)

	log.Fatal(s.ListenAndServe())
}
