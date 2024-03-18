package model

type KlineRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	TimeZone  string `json:"timeZone"`
	Limit     int    `json:"limit"`
}
