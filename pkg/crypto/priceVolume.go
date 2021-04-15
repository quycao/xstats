package crypto

import (
	"fmt"
	"time"
)

// PriceVolume is struct of price volume klines data
type PriceVolume struct {
	OpenTime         time.Time `json:"openTime"`
	OpenPrice        float64   `json:"openPrice"`
	HighPrice        float64   `json:"highPrice"`
	LowPrice         float64   `json:"lowPrice"`
	ClosePrice       float64   `json:"closePrice"`
	Volume           float64   `json:"volume"`
	CloseTime        time.Time `json:"closeTime"`
	QuoteAssetVolume float64   `json:"quoteAssetVolume"`
	NumberOfTrades   int64     `json:"numberOfTrades"`
	TakerBuyBase     float64   `json:"takerBuyBase"`
	TakerBuyQuote    float64   `json:"takerBuyQuote"`
	Ignore           float64   `json:"Ignore"`
}

// PriceVolumeStatsResult is result of price volume stats
type PriceVolumeStatsResult struct {
	Time                      time.Time `json:"time"`
	Symbol                    string    `json:"symbol"`
	Period                    string    `json:"period"`
	Price                     float64   `json:"price"`
	Volume                    float64   `json:"volume"`
	AvgVolume10Periods        float64   `json:"avgVolume10Periods"`
	HighestPrice30Periods     float64   `json:"highestPrice30Periods"`
	RatioChangeVol10Periods   float64   `json:"ratioChangeVol10Days"`
	RatioChangePrice30Periods float64   `json:"ratioChangePrice30Days"`
	RatioChangePrice          float64   `json:"ratioChangePrice"`
	Suggestion                string    `json:"suggestion"`
}

func (r *PriceVolumeStatsResult) ToString() string {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	str := fmt.Sprintf(`
		Time: %s
		Symbol: %s
		Period: %s
		Price: %f
		Volume: %f
		Avg Volume: %f
		Highest Price: %f
		Change Volume 10 periods: %.4f%%
		Change Price 30 periods: %.4f%%
		Change Price this period: %.4f%%
		Suggestion: %s
		`,
		r.Time.In(loc).Format("2006-01-02 15:04"),
		r.Symbol,
		r.Period,
		r.Price,
		r.Volume,
		r.AvgVolume10Periods,
		r.HighestPrice30Periods,
		r.RatioChangeVol10Periods*100,
		r.RatioChangePrice30Periods*100,
		r.RatioChangePrice*100,
		r.Suggestion,
	)

	return str
}
