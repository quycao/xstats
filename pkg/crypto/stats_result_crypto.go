package crypto

import (
	"fmt"
	"math"
	"time"
)

// StatsResultCrypto is result of statistic
type StatsResultCrypto struct {
	Time       time.Time `json:"time"`
	Symbol     string    `json:"symbol"`
	AvgPrice   float64   `json:"avg_price"`
	BuyVol     float64   `json:"buy_vol"`
	SellVol    float64   `json:"sell_vol"`
	BuySellPct int64     `json:"buy_sell_pct"`
	Status     string    `json:"status"`
	Suggestion string    `json:"suggestion"`
}

func (r *StatsResultCrypto) ToString() string {
	// return fmt.Sprintf("Symbol:     %12v | Avg Price:  %12v | Buy Vol:    %12v | Sell Vol:   %12v | Buy/Sell:  %12v%% | Status:     %12v | Suggestion: %12v\n",
	// 	r.Symbol,
	// 	math.Round(r.AvgPrice*100)/100,
	// 	math.Round(r.BuyVol*100)/100,
	// 	math.Round(r.SellVol*100)/100,
	// 	r.BuySellPct,
	// 	r.Status,
	// 	r.Suggestion)

	return fmt.Sprintf("%s | %s | Avg Price: %9v | Buy Vol: %8v | Sell Vol: %8v | Buy/Sell: %5v%% | Status: %11v | Suggestion: %5v\n",
		r.Time.Format("15:04"),
		r.Symbol,
		math.Round(r.AvgPrice*100)/100,
		math.Round(r.BuyVol*100)/100,
		math.Round(r.SellVol*100)/100,
		r.BuySellPct,
		r.Status,
		r.Suggestion)
}
