package rest

import (
	"context"
	v1 "github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterInfoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatal("failed to start HTTP gateway", errors.Wrap(err, "Register Endpoint"))
	}

	srv := &http.Server{
		Addr: ":" + httpPort,
		// add handler with middleware
		//Handler: middleware.AddRequestID(
		//	middleware.AddLogger(logger.Log, mux)),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Printf("starting HTTP/REST gateway...\n")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", mux)
	//return srv.ListenAndServe()

	return nil
}
