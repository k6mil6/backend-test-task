package klines

import (
	"fmt"
	"strconv"
)

type Response struct {
	OpenTime                 int64   `json:"openTime"`
	OpenPrice                float64 `json:"open_price"`
	HighPrice                float64 `json:"high_price"`
	LowPrice                 float64 `json:"low_price"`
	ClosePrice               float64 `json:"close_price"`
	Volume                   float64 `json:"volume"`
	CloseTime                int64   `json:"kline_close_time"`
	QuoteAssetVolume         float64 `json:"quote_asset_volume"`
	NumberOfTrades           int64   `json:"number_of_trades"`
	TakerBuyBaseAssetVolume  float64 `json:"taker_buy_base_asset_volume"`
	TakerBuyQuoteAssetVolume float64 `json:"taker_buy_quote_asset_volume"`
	Ignore                   int64   `json:"ignore,omitempty"`
}

type KlineResponse []interface{}

func getFloat64(raw interface{}) float64 {
	switch v := raw.(type) {
	case float64:
		return v
	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return f
	default:
		return 0
	}
}

func getInt64(raw interface{}) int64 {
	switch v := raw.(type) {
	case float64:
		return int64(v)
	case string:
		i, _ := strconv.ParseInt(v, 10, 64)
		return i
	default:
		return 0
	}
}

func parseKline(kline KlineResponse) (Response, error) {
	if len(kline) < 11 {
		return Response{}, fmt.Errorf("invalid kline response length")
	}

	return Response{
		OpenTime:                 getInt64(kline[0]),
		OpenPrice:                getFloat64(kline[1]),
		HighPrice:                getFloat64(kline[2]),
		LowPrice:                 getFloat64(kline[3]),
		ClosePrice:               getFloat64(kline[4]),
		Volume:                   getFloat64(kline[5]),
		CloseTime:                getInt64(kline[6]),
		QuoteAssetVolume:         getFloat64(kline[7]),
		NumberOfTrades:           getInt64(kline[8]),
		TakerBuyBaseAssetVolume:  getFloat64(kline[9]),
		TakerBuyQuoteAssetVolume: getFloat64(kline[10]),
		Ignore:                   getInt64(kline[11]),
	}, nil
}
