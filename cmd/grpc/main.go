package main

import (
	"github.com/k6mil6/backend-test-task/internal/grpcservice/app"
	"github.com/k6mil6/backend-test-task/pkg/config"
	"github.com/k6mil6/backend-test-task/pkg/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig(".")

	log := logger.SetupLogger(cfg.Env).With(slog.String("env", cfg.Env))
	log.Debug("logger debug mode enabled")

	application := app.New(log, cfg.GrpcPort)

	go func() {
		err := application.GRPCServer.Run()
		if err != nil {
			log.Error("failed to start server", err)
			return
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	stopSignal := <-stop

	log.Info("shutting down server", slog.String("signal", stopSignal.String()))

	application.GRPCServer.Stop()

	log.Info("server stopped")
}
