package v1

import (
	"context"
	"github.com/3hajk/grpc-http-rest-microservice/app/store"
	"testing"
	"time"

	"github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
)

func Test_InfoServiceServer_Info(t *testing.T) {
	ctx := context.Background()

	info := store.NewInfo(ctx, 5*time.Minute)
	s := NewInfoServiceServer(info)

	type args struct {
		//nolin:containedctx
		ctx context.Context
		req *v1.InfoRequest
	}
	tests := []struct {
		name    string
		s       v1.InfoServiceServer
		args    args
		mock    func()
		want    *v1.InfoResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.InfoRequest{
					Api: "v1",
				},
			},
			mock: func() {},
			want: &v1.InfoResponse{
				Api:  "v1",
				Info: &v1.Info{},
			},
		},
		{
			name: "Unsupported API",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.InfoRequest{
					Api: "v1000",
				},
			},
			mock:    func() {},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Info(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("toDoServiceServer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
