package main

import (
	"github.com/fasthttp/router"
	"github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/handlers/user"
	"github.com/k6mil6/backend-test-task/pkg/config"
	"github.com/k6mil6/backend-test-task/pkg/logger"
	"github.com/valyala/fasthttp"
	"log/slog"
)

func main() {
	cfg := config.LoadConfig(".")

	log := logger.SetupLogger(cfg.Env).With(slog.String("env", cfg.Env))
	log.Debug("logger debug mode enabled")

	r := router.New()

	r.GET("/klines", user.GetKlines)

	err := fasthttp.ListenAndServe(cfg.HttpAddress, r.Handler)
	if err != nil {
		log.Error("failed to start server", err)
	}
}
