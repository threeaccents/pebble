package grpc

import (
	"context"
	"net"
	"time"

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

func (s *Server) SetTTL(ctx context.Context, req *pb.SetTTLRequest) (*pb.SetResponse, error) {
	ttl := time.Duration(req.Ttl) * time.Second

	if err := s.Storage.SetTTL(req.Key, req.Value, ttl); err != nil {
		return nil, err
	}

	return &pb.SetResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if err := s.Storage.Delete(req.Key); err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{}, nil
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
