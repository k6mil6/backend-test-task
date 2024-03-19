package main

import (
	"context"
	"github.com/fasthttp/router"
	grpcclient "github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/clients/grpc"
	"github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/handlers/user"
	"github.com/k6mil6/backend-test-task/internal/storage/postgres"
	"github.com/k6mil6/backend-test-task/pkg/config"
	"github.com/k6mil6/backend-test-task/pkg/logger"
	"github.com/valyala/fasthttp"
	"log/slog"
	"time"
)

func main() {
	cfg := config.LoadConfig(".")

	log := logger.SetupLogger(cfg.Env).With(slog.String("env", cfg.Env))
	log.Debug("logger debug mode enabled")

	ctx := context.Background()

	grpcClient, err := grpcclient.New(ctx, log, cfg.ClientAddress, cfg.ClientTimeout, cfg.ClientRetries)
	if err != nil {
		log.Error("failed to create grpc client", err)
		return
	}

	storage, err := postgres.New(cfg.DbAddress, 3, 5*time.Second)
	if err != nil {
		log.Error("failed to create storage", err)
		return
	}

	r := router.New()

	r.GET("/klines", user.GetKlines(ctx, grpcClient, storage))

	err = fasthttp.ListenAndServe(cfg.HttpAddress, r.Handler)
	if err != nil {
		log.Error("failed to start server", err)
	}
}
