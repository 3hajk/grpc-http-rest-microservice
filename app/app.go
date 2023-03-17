package app

import (
	"context"
	"fmt"
	"log"

	"github.com/3hajk/grpc-http-rest-microservice/app/protocol/grpc"
	"github.com/3hajk/grpc-http-rest-microservice/app/protocol/rest"
	"github.com/3hajk/grpc-http-rest-microservice/app/service/v1"
	"github.com/3hajk/grpc-http-rest-microservice/app/store"
	"github.com/3hajk/grpc-http-rest-microservice/cfg"
	"github.com/pkg/errors"
)

var (
	Version string
	Build   string
	Branch  string
)

func GetAppTitle() string {
	return fmt.Sprintf("Matcher service (%s) Version: %s, Build Time: %s", Branch, Version, Build)
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	conf, err := cfg.Read()
	if err != nil {
		return errors.Wrap(err, "read config")
	}

	log.Printf("%+v", conf)

	ctx := context.Background()

	info := store.NewInfo(ctx, conf.Regenerate)

	v1API := v1.NewInfoServiceServer(info)

	go func() {
		err = rest.RunServer(ctx, conf.GRPCService.Port, conf.HTTPService.Port)
		if err != nil {
			log.Fatalf("cant start HTTP service: %v", err)
		}
	}()

	return grpc.RunServer(ctx, v1API, conf.GRPCService.Port)
}
