package model

import "fmt"

// StatsResultTCBS is result of statistic
type StatsResultTCBS struct {
	Ticker     string `json:"ticker"`
	BuyVol     int64  `json:"buy_vol"`
	SellVol    int64  `json:"sell_vol"`
	BuySellPct int64  `json:"buy_sell_pct"`
	Status     string `json:"status"`
	Suggestion string `json:"suggestion"`
}

func (r *StatsResultTCBS) ToString() string {
	return fmt.Sprintf("\nTicker:     %12v\nBuy Vol:    %12v\nSell Vol:   %12v\nBuy/Sell:  %12v%%\nStatus:     %12v\nSuggestion: %12v\n",
		r.Ticker, r.BuyVol, r.SellVol, (r.BuyVol-r.SellVol)*100/r.SellVol, r.Status, r.Suggestion)
}
