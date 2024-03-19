package app

import (
	grpcapp "github.com/k6mil6/backend-test-task/internal/grpcservice/app/grpc"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
) *App {

	grpcServer := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: grpcServer,
	}
}
