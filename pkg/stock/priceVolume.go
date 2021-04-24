package stock

import (
	"fmt"
)

// PriceVolume is data of stock price and volume
type PriceVolume struct {
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Close       float64 `json:"close"`
	Volume      int64   `json:"volume"`
	TradingDate string  `json:"tradingDate"`
}

// CurrentPriceVolume is current price data
type CurrentPriceVolume struct {
	Ticker            string  `json:"Ticker"`
	CeilingPrice      float64 `json:"ceilingPrice"`
	FloorPrice        float64 `json:"floorPrice"`
	OpenPrice         float64 `json:"openPrice"`
	HighPrice         float64 `json:"highPrice"`
	LowPrice          float64 `json:"lowPrice"`
	CurrentPrice      float64 `json:"currentPrice"`
	ReferencePrice    float64 `json:"referencePrice"`
	AveragePrice      float64 `json:"averagePrice"`
	Change            float64 `json:"change"`
	AccumulatedVolume int64   `json:"accumulatedVolume"`
	AccumulatedValue  float64 `json:"accumulatedValue"`
	TradingDate       string  `json:"tradingDate"`
	// ForeignRemain         int64   `json:"foreignRemain"`
	// ForeignRoom           int64   `json:"foreignRoom"`
	// Bid1Price             float64 `json:"bid1Price"`
	// Bid2Price             float64 `json:"bid2Price"`
	// Bid3Price             float64 `json:"bid3Price"`
	// Ask1Price             float64 `json:"ask1Price"`
	// Ask2Price             float64 `json:"ask2Price"`
	// Ask3Price             float64 `json:"ask3Price"`
	// Bid1Volume            int64   `json:"bid1Volume"`
	// Bid2Volume            int64   `json:"bid2Volume"`
	// Bid3Volume            int64   `json:"bid3Volume"`
	// Ask1Volume            int64   `json:"ask1Volume"`
	// Ask2Volume            int64   `json:"ask2Volume"`
	// Ask3Volume            int64   `json:"ask3Volume"`
	// ForeignBuy            int64   `json:"foreignBuy"`
	// ForeignSell           int64   `json:"foreignSell"`
	// PtTotalTradedQuantity int64   `json:"ptTotalTradedQuantity"`
	// PtTotalTradedValue    float64 `json:"ptTotalTradedValue"`
	// TotalAskQuantity      int64   `json:"totalAskQuantity"`
	// TotalBidQuantity      int64   `json:"totalBidQuantity"`
}

// PriceVolume2 is data of stock price and volume
type PriceVolume2 struct {
	Price            int64   `json:"p"`
	ChangePrice      int64   `json:"cp"`
	RatioChangePrice float64 `json:"rcp"`
	Volume           int64   `json:"v"`
	Date             string  `json:"dt"`
}

// PriceVolumeStatsResult is result of price volume stats
type PriceVolumeStatsResult struct {
	Ticker                 string  `json:"ticker"`
	Price                  float64 `json:"price"`
	Volume                 int64   `json:"volume"`
	Trend10Days            string  `json:"trend"`
	AvgVolume10Days        int64   `json:"avg_volume_10_days"`
	HighestPrice30Days     float64 `json:"highest_price_30_days"`
	RatioChangeVol10Days   float64 `json:"ratio_change_vol_10_days"`
	RatioChangePrice30Days float64 `json:"ratio_change_price_30_days"`
	RatioChangePrice       float64 `json:"ratio_change_price"`
	Date                   string  `json:"Date"`
	Suggestion             string  `json:"suggestion"`
	Reason                 string  `json:"reason"`
}

// PriceVolumeBind is binding for PriceVolume
type PriceVolumeBind struct {
	Ticker string         `json:"ticker"`
	Data   []*PriceVolume `json:"data"`
}

func (r *PriceVolumeStatsResult) ToString() string {
	str := fmt.Sprintf(`
	Date: %s
	Symbol: %s
	Price: %.0f
	Trend: %s
	Volume: %d
	Avg Volume 10 days: %d
	Highest Price 30 days: %.0f
	Change Volume 10 days: %.2f%%
	Change Price 30 days: %.2f%%
	Change Price today: %.2f%%
	Suggestion: %s
	Reason: %s
	`,
		r.Date,
		r.Ticker,
		r.Price,
		r.Trend10Days,
		r.Volume,
		r.AvgVolume10Days,
		r.HighestPrice30Days,
		r.RatioChangeVol10Days*100,
		r.RatioChangePrice30Days*100,
		r.RatioChangePrice*100,
		r.Suggestion,
		r.Reason,
	)

	return str
}
