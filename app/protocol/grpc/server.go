package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish Info service
func RunServer(ctx context.Context, v1API v1.InfoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterInfoServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")

	return server.Serve(listen)
}
