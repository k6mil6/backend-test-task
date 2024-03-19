package postgres

import (
	servermodel "github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/model"
	"github.com/k6mil6/backend-test-task/internal/storage/postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	*gorm.DB
}

func New(dsn string, maxRetries int, delay time.Duration) (*Storage, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
		if err != nil {
			if i < maxRetries-1 {
				time.Sleep(delay)
			}
		}
	}

	if err := db.AutoMigrate(&model.KlineResponse{}); err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

func (s *Storage) AddKlines(symbol, interval string, klines []servermodel.KlineResponse) error {
	dbKlines := make([]model.KlineResponse, len(klines))

	for i, kline := range klines {
		dbKlines[i].Symbol = symbol
		dbKlines[i].Interval = interval
		dbKlines[i].OpenTime = kline.OpenTime
		dbKlines[i].OpenPrice = kline.OpenPrice
		dbKlines[i].HighPrice = kline.HighPrice
		dbKlines[i].LowPrice = kline.LowPrice
		dbKlines[i].ClosePrice = kline.ClosePrice
		dbKlines[i].Volume = kline.Volume
		dbKlines[i].CloseTime = kline.CloseTime
		dbKlines[i].QuoteAssetVolume = kline.QuoteAssetVolume
		dbKlines[i].NumberOfTrades = kline.NumberOfTrades
		dbKlines[i].TakerBuyBaseAssetVolume = kline.TakerBuyBaseAssetVolume
		dbKlines[i].TakerBuyQuoteAssetVolume = kline.TakerBuyQuoteAssetVolume
		dbKlines[i].Ignore = kline.Ignore
	}

	return s.Create(&dbKlines).Error
}
