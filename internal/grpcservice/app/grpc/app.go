package grpcapp

import (
	"google.golang.org/grpc"
	"log/slog"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
}
