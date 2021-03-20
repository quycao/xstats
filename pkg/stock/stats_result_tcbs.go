package stock

import (
	"fmt"
	"time"
)

// StatsResultTCBS is result of statistic
type StatsResultTCBS struct {
	Time        time.Time `json:"time"`
	Ticker      string    `json:"ticker"`
	AvgPrice    int64     `json:"avg_price"`
	ChangePrice int64     `json:"change_price"`
	BuyVol      int64     `json:"buy_vol"`
	SellVol     int64     `json:"sell_vol"`
	BuySellPct  int64     `json:"buy_sell_pct"`
	Status      string    `json:"status"`
	Suggestion  string    `json:"suggestion"`
}

func (r *StatsResultTCBS) ToString() string {
	return fmt.Sprintf("\nTime:       %12v\nTicker:     %12v\nAvg Price:  %12v\nChg Price:  %12v\nBuy Vol:    %12v\nSell Vol:   %12v\nBuy/Sell:  %12v%%\nStatus:     %12v\nSuggestion: %12v\n",
		r.Time, r.Ticker, r.AvgPrice, r.ChangePrice, r.BuyVol, r.SellVol, r.BuySellPct, r.Status, r.Suggestion)
}
