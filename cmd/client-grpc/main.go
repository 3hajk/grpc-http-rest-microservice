package client_grpc

import (
	"context"
	"flag"
	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var (
	address = flag.String("server", "", "gRPC server in format host:port")
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

	//t := time.Now().In(time.UTC)
	//reminder, _ := ptypes.TimestampProto(t)
	//pfx := t.Format(time.RFC3339Nano)

	// Call Info
	req := v1.InfoRequest{
		Api: apiVersion,
		//: &v1.ToDo{
		//	Title:       "title (" + pfx + ")",
		//	Description: "description (" + pfx + ")",
		//	Reminder:    reminder,
		//},
	}
	res, err := c.Info(ctx, &req)
	if err != nil {
		log.Fatalf("Info request failed: %v", err)
	}
	log.Printf("Info result: <%+v>\n", res)
}
