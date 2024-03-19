package model

import "gorm.io/gorm"

type KlineResponse struct {
	gorm.Model
	Symbol                   string `gorm:"index:idx_symbol"`
	Interval                 string `gorm:"index:idx_interval"`
	OpenTime                 int64
	OpenPrice                float64
	HighPrice                float64
	LowPrice                 float64
	ClosePrice               float64
	Volume                   float64
	CloseTime                int64
	QuoteAssetVolume         float64
	NumberOfTrades           int64
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   int64
}
