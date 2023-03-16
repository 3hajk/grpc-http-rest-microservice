package main

import (
	"context"
	"flag"
	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var (
	address = flag.String("server", ":9090", "gRPC server in format host:port")
)

func main() {
	// get configuration
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewInfoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Info
	req := v1.InfoRequest{
		Api: apiVersion,
	}
	res, err := c.Info(ctx, &req)
	if err != nil {
		log.Fatalf("Info request failed: %v", err)
	}
	log.Printf("Info result: <%+v>\n", res)
}
