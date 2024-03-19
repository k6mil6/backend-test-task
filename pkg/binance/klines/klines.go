package klines

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

const url = "https://api.binance.com/api/v3/klines"

type GetKlinesOptions struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	TimeZone  string `json:"time_zone,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

func Get(symbol, interval string, opts ...func(*GetKlinesOptions)) ([]Response, error) {
	config := GetKlinesOptions{
		Symbol:   symbol,
		Interval: interval,
	}

	for _, opt := range opts {
		opt(&config)
	}

	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)
	args.Add("symbol", config.Symbol)
	args.Add("interval", config.Interval)

	if config.StartTime != 0 {
		args.Add("startTime", strconv.FormatInt(config.StartTime, 10))
	}
	if config.EndTime != 0 {
		args.Add("endTime", strconv.FormatInt(config.EndTime, 10))
	}
	if config.Limit != 0 {
		args.Add("limit", strconv.Itoa(config.Limit))
	}
	if config.TimeZone != "" {
		args.Add("timeZone", config.TimeZone)
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod("GET")
	req.SetRequestURI(url)
	req.URI().SetQueryStringBytes(args.QueryString())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		if errors.Is(err, fasthttp.ErrTimeout) {
			return nil, status.Error(codes.DeadlineExceeded, "request timed out")
		}
		return nil, status.Error(codes.Internal, "error getting klines")
	}

	bodyBytes := resp.Body()

	var rawKlines []KlineResponse
	if err := json.Unmarshal(bodyBytes, &rawKlines); err != nil {
		return nil, status.Error(codes.Internal, "error unmarshalling klines")
	}

	var klines []Response
	for _, rawKline := range rawKlines {
		kline, err := parseKline(rawKline)
		if err != nil {
			return nil, status.Error(codes.Internal, "error parsing kline")
		}
		klines = append(klines, kline)
	}

	return klines, nil
}

func WithStartTime(startTime int64) func(*GetKlinesOptions) {
	return func(o *GetKlinesOptions) {
		o.StartTime = startTime
	}
}

func WithEndTime(endTime int64) func(*GetKlinesOptions) {
	return func(o *GetKlinesOptions) {
		o.EndTime = endTime
	}
}

func WithTimeZone(timeZone string) func(*GetKlinesOptions) {
	return func(o *GetKlinesOptions) {
		o.TimeZone = timeZone
	}
}

func WithLimit(limit int) func(*GetKlinesOptions) {
	return func(o *GetKlinesOptions) {
		o.Limit = limit
	}
}
