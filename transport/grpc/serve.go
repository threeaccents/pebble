package grpc

import (
	"github.com/threeaccents/cache"
	"github.com/threeaccents/cache/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type ServeOptions struct {
	Port string
}

func ServePort(port string) func(*ServeOptions) {
	return func(opt *ServeOptions) {
		opt.Port = port
	}
}

func Serve(storage cache.Storage, opts ...func(*ServeOptions)) error {
	options := &ServeOptions{
		Port: ":4200",
	}

	for _, option := range opts {
		option(options)
	}

	listen, err := net.Listen("tcp", options.Port)
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
