package main

import (
	"flag"
	"log"
	"os"

	"github.com/oriiolabs/pebble/adapter/badger"
	"github.com/rs/zerolog"

	"github.com/oriiolabs/pebble/transport/grpc"
)

var (
	dbDirPtr = flag.String("dbdir", "./db", "the directory were the cache is stored")
	portPtr  = flag.String("port", ":4210", "the port were to run the grpc server")
)

func main() {
	flag.Parse()

	logger := zerolog.New(os.Stdout).With().Timestamp().Str("app", "oriio").Logger()
	logger.Info().Msg("Starting Service...")

	db, err := badger.Open(*dbDirPtr)
	if err != nil {
		panic(err)
	}

	storage := &badger.CacheStorage{
		DB:  db,
		Log: logger,
	}

	s := grpc.NewServer(
		storage,
		grpc.ServerPort(*portPtr),
	)

	logger.Info().Str("port", *portPtr).Msg("listening")

	log.Fatal(s.ListenAndServe())
}
