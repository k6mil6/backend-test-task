package grpc

import (
	"context"
	"fmt"
	"github.com/k6mil6/backend-test-task/pkg/binance/klines"
	grpcv1 "github.com/k6mil6/protobuf/gen/go/grpc"
	"google.golang.org/grpc"
)

type serverApi struct {
	grpcv1.UnimplementedGrpcServiceServer
}

func Register(gRPC *grpc.Server) {
	grpcv1.RegisterGrpcServiceServer(gRPC, &serverApi{})
}

func (s *serverApi) GetKlines(
	ctx context.Context,
	req *grpcv1.GetKlinesRequest,
) (*grpcv1.GetKlinesResponse, error) {

	symbol := req.Symbol
	interval := req.Interval
	startTime := req.StartTime
	endTime := req.EndTime
	timeZone := req.TimeZone
	limit := req.Limit

	fmt.Println(symbol, interval, startTime, endTime, limit)

	var opts []func(*klines.GetKlinesOptions)
	if startTime != 0 {
		opts = append(opts, klines.WithStartTime(startTime))
	}
	if endTime != 0 {
		opts = append(opts, klines.WithEndTime(endTime))
	}
	if limit != 0 {
		opts = append(opts, klines.WithLimit(int(limit)))
	}
	if timeZone != "" {
		opts = append(opts, klines.WithTimeZone(timeZone))
	}

	allKlines, err := klines.Get(symbol, interval, opts...)
	if err != nil {
		return nil, err
	}

	grpcResponse := &grpcv1.GetKlinesResponse{
		Klines: make([]*grpcv1.Kline, len(allKlines)),
	}
	for i, kline := range allKlines {
		grpcResponse.Klines[i] = &grpcv1.Kline{
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

	return grpcResponse, nil
}
