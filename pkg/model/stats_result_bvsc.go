package model

// StatsResultBVSC is result of statistic
type StatsResultBVSC struct {
	Symbol     string `json:"symbol"`
	BuyVol     int64  `json:"buy_vol"`
	SellVol    int64  `json:"sell_vol"`
	Status     string `json:"status"`
	Suggestion string `json:"suggestion"`
}
