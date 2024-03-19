package user

import (
	"context"
	"encoding/json"
	"github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/model"
	"github.com/valyala/fasthttp"
	"strconv"
)

type KlinesProvider interface {
	GetKlines(ctx context.Context, request model.KlineRequest) ([]model.KlineResponse, error)
}

type KlinesSaver interface {
	AddKlines(symbol, interval string, klines []model.KlineResponse) error
}

func GetKlines(context context.Context, provider KlinesProvider, saver KlinesSaver) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var req model.KlineRequest
		req.Symbol = string(ctx.QueryArgs().Peek("symbol"))
		req.Interval = string(ctx.QueryArgs().Peek("interval"))
		req.StartTime = parseOptionalInt64(ctx, "start_time")
		req.EndTime = parseOptionalInt64(ctx, "end_time")
		req.TimeZone = string(ctx.QueryArgs().Peek("time_zone"))
		req.Limit = parseOptionalInt(ctx, "limit")

		if req.Symbol == "" || req.Interval == "" {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.SetBodyString("Missing required query parameters: symbol or interval")
			return
		}

		resp, err := provider.GetKlines(context, req)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		if err := saver.AddKlines(req.Symbol, req.Interval, resp); err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetBody(jsonResp)
		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}

func parseOptionalInt64(ctx *fasthttp.RequestCtx, paramName string) int64 {
	paramValue := ctx.QueryArgs().Peek(paramName)
	if paramValue != nil {
		value, err := strconv.ParseInt(string(paramValue), 10, 64)
		if err == nil {
			return value
		}
	}
	return 0
}

func parseOptionalInt(ctx *fasthttp.RequestCtx, paramName string) int {
	paramValue := ctx.QueryArgs().Peek(paramName)
	if paramValue != nil {
		value, err := strconv.Atoi(string(paramValue))
		if err == nil {
			return value
		}
	}
	return 0
}
