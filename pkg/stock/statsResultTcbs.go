package stock

import (
	"fmt"
	"math"
	"time"
)

// StatsResult is result of statistic
type StatsResult struct {
	Time          time.Time `json:"time"`
	Ticker        string    `json:"ticker"`
	BuySellActive float64   `json:"buySellActive"`
	BidAskRatio   float64   `json:"bidAskRatio"`
	Volumn        int64     `json:"volumn"`
	Status        string    `json:"status"`
	Suggestion    string    `json:"suggestion"`
}

func (r *StatsResult) ToString() string {
	return fmt.Sprintf("\nTime:       %12s\nTicker:     %12v\nB/S Active: %12v\nBid/Ask:    %12v\nVolumn:     %12v\nStatus:     %12v\nSuggestion: %12v\n",
		r.Time.Format("2006-01-02 15:04"), r.Ticker, r.BuySellActive, math.Floor(r.BidAskRatio*1000)/1000, r.Volumn, r.Status, r.Suggestion)
}
