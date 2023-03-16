package v1

import (
	"context"
	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type infoServiceServer struct {
	UUID string
}

// NewToDoServiceServer creates Info service
func NewInfoServiceServer() v1.InfoServiceServer {
	return &infoServiceServer{UUID: ""}
}
func (s *infoServiceServer) Info(ctx context.Context, request *v1.InfoRequest) (*v1.InfoResponse, error) {
	panic("implement me")
}

func (s *infoServiceServer) mustEmbedUnimplementedInfoServiceServer() {
	panic("implement me")
}
