package grpc

import (
	"net"

	"github.com/oriiolabs/pebble"

	"github.com/oriiolabs/pebble/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const defaultPort = ":4200"

type Server struct {
	Storage pebble.Storage
	Port    string
}

func ServerPort(port string) func(*Server) {
	return func(opt *Server) {
		opt.Port = port
	}
}

func NewServer(storage pebble.Storage, opts ...func(*Server)) *Server {
	s := &Server{
		Storage: storage,
		Port:    defaultPort,
	}

	for _, option := range opts {
		option(s)
	}

	return s
}

func (s *Server) ListenAndServe() error {
	listen, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCacheServer(grpcServer, s)

	reflection.Register(grpcServer)

	return grpcServer.Serve(listen)
}
