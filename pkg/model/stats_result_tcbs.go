package model

// StatsResultTCBS is result of statistic
type StatsResultTCBS struct {
	Ticker     string `json:"ticker"`
	BuyVol     int64  `json:"buy_vol"`
	SellVol    int64  `json:"sell_vol"`
	Status     string `json:"status"`
	Suggestion string `json:"suggestion"`
}
