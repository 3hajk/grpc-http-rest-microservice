package v1

import (
	"context"

	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"github.com/3hajk/grpc-http-rest-microservice/app/store"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// infoServiceServer is implementation of v1.InfoServiceServer proto interface
type infoServiceServer struct {
	info *store.Info
}

// NewInfoServiceServer creates Info service
//nolint:ireturn
func NewInfoServiceServer(storeInfo *store.Info) v1.InfoServiceServer {
	return &infoServiceServer{
		info: storeInfo,
	}
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

func (i *infoServiceServer) Info(ctx context.Context, req *v1.InfoRequest) (*v1.InfoResponse, error) {
	// check if the API version requested by client is supported by server
	if err := i.checkAPI(req.Api); err != nil {
		return nil, err
	}

	UUID, hash, generationTime := i.info.GetInfo()

	t, err := ptypes.TimestampProto(generationTime)
	if err != nil {
		return nil, err
	}
	return &v1.InfoResponse{
		Api: apiVersion,
		Info: &v1.Info{
			Uuid:           UUID,
			Hash:           hash,
			GenerationTime: t,
		},
	}, nil
}
