package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterInfoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		return errors.WithMessage(errors.Wrap(err, "Register Endpoint"), "failed to start HTTP gateway")
	}

	log.Printf("starting HTTP/REST gateway...\n")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	// nolint:gosec
	return http.ListenAndServe(":"+httpPort, mux)
}
