package stock

// PriceVolume is data of stock price and volume
type PriceVolume struct {
	Price            int64   `json:"p"`
	ChangePrice      int64   `json:"cp"`
	RatioChangePrice float64 `json:"rcp"`
	Volume           int64   `json:"v"`
	Date             string  `json:"dt"`
}

// PriceVolumeBind is binding for PriceVolume
type PriceVolumeBind struct {
	Ticker string         `json:"ticker"`
	Data   []*PriceVolume `json:"data"`
}
