package grpc

import (
	"context"
	"fmt"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/model"
	grpcv1 "github.com/k6mil6/protobuf/gen/go/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"
)

type Client struct {
	api grpcv1.GrpcServiceClient
	log *slog.Logger
}

func New(
	ctx context.Context,
	log *slog.Logger,
	address string,
	timeout time.Duration,
	retriesCount int,
) (*Client, error) {
	op := "grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadSent, grpclog.PayloadReceived),
	}

	cc, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Client{
		api: grpcv1.NewGrpcServiceClient(cc),
	}, nil
}

func (c *Client) GetKlines(ctx context.Context, request model.KlineRequest) ([]model.KlineResponse, error) {
	const op = "grpc.Client.GetKlines"

	response, err := c.api.GetKlines(ctx, &grpcv1.GetKlinesRequest{
		Symbol:    request.Symbol,
		Interval:  request.Interval,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		TimeZone:  request.TimeZone,
		Limit:     int32(request.Limit),
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	klinesResponses := make([]model.KlineResponse, len(response.Klines))

	for i, kline := range response.Klines {
		klinesResponses[i] = model.KlineResponse{
			OpenTime:                 kline.OpenTime,
			OpenPrice:                kline.OpenPrice,
			HighPrice:                kline.HighPrice,
			LowPrice:                 kline.LowPrice,
			ClosePrice:               kline.ClosePrice,
			Volume:                   kline.Volume,
			CloseTime:                kline.CloseTime,
			QuoteAssetVolume:         kline.QuoteAssetVolume,
			NumberOfTrades:           kline.NumberOfTrades,
			TakerBuyBaseAssetVolume:  kline.TakerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: kline.TakerBuyQuoteAssetVolume,
		}
	}

	return klinesResponses, nil
}

func InterceptorLogger(log *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, level grpclog.Level, msg string, fields ...any) {
		log.Log(ctx, slog.Level(level), msg, fields...)
	})
}
