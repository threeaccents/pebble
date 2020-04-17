package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	"github.com/threeaccents/cache"
	"github.com/threeaccents/cache/pb"
)

type Server struct {
	Storage cache.Storage
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	value, err := s.Storage.Get(req.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Value: value,
	}, nil
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	if err := s.Storage.Set(req.Key, req.Value); err != nil {
		return nil, err
	}

	return &pb.SetResponse{}, nil
}

func Serve(storage cache.Storage) error {
	listen, err := net.Listen("tcp", ":4200")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCacheServer(grpcServer, &Server{
		Storage: storage,
	})

	reflection.Register(grpcServer)

	return grpcServer.Serve(listen)
}
