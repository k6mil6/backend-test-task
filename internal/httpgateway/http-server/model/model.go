package model

type KlineRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	TimeZone  string `json:"time_zone,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

type KlineResponse struct {
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
