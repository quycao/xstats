package stock

// BidAsk is Avg Over Buy Percent
type BidAsk struct {
	Obp  float64 `json:"obp"`
	Osp  float64 `json:"osp"`
	Aobp float64 `json:"aobp"`
	Time string  `json:"t"`
}

// BidAskPercentByDay is BidAskPercent by day
type BidAskPercentByDay struct {
	Ticker    string  `json:"ticker"`
	OBPercent float64 `json:"obPercent"`
	Date      string  `json:"d"`
}

// BidAskBind is binding for BidAsk
type BidAskBind struct {
	Ticker        string    `json:"ticker"`
	OverBidAskLog []*BidAsk `json:"overBidAskLog"`
	Date          string    `json:"d"`
}
