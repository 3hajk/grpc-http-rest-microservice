package v1

import (
	"context"
	v1 "github.com/3hajk/grpc-http-rest-microservice/app/api/v1"
	"testing"
)

func Test_toDoServiceServer_Create(t *testing.T) {
	ctx := context.Background()

	s := NewInfoServiceServer()
	//tm := time.Now().In(time.UTC)
	//reminder, _ := ptypes.TimestampProto(tm)

	type args struct {
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
