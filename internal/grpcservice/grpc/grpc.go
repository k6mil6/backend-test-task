package grpc

import (
	"context"
	grpcv1 "github.com/k6mil6/protobuf/gen/go/grpc"
	"google.golang.org/grpc"
)

type serverApi struct {
	grpcv1.UnimplementedGrpcServiceServer
}

func Register(gRPC *grpc.Server) {
	grpcv1.RegisterGrpcServiceServer(gRPC, &serverApi{})
}

func (*serverApi) GetKlines(
	ctx context.Context,
	req *grpcv1.GetKlinesRequest,
) (*grpcv1.GetKlinesResponse, error) {
	return nil, nil
}
