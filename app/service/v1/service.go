package v1

import (
	"context"
	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// infoServiceServer is implementation of v1.InfoServiceServer proto interface
type infoServiceServer struct {
	UUID     string
	Reminder *timestamp.Timestamp
}

// NewInfoServiceServer creates Info service
func NewInfoServiceServer() v1.InfoServiceServer {
	return &infoServiceServer{}
}

func (i *infoServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (i *infoServiceServer) generateUUID() error {

	i.UUID = "werwerwerwe"
	t := time.Now()
	reminder, err := ptypes.TimestampProto(t)

	if err != nil {
		return err
	}
	i.Reminder = reminder
	return nil
}

func (i *infoServiceServer) Info(ctx context.Context, req *v1.InfoRequest) (*v1.InfoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := i.checkAPI(req.Api); err != nil {
		return nil, err
	}
	return &v1.InfoResponse{
		Api: apiVersion,
		Info: &v1.Info{
			Id:          12345,
			Title:       i.UUID,
			Description: "test",
			Reminder:    i.Reminder,
		},
	}, nil
}

func (i *infoServiceServer) mustEmbedUnimplementedInfoServiceServer() {
	log.Println("mustEmbedUnimplementedInfoServiceServer")
}
