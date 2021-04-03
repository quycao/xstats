package stock

// BSA is Buy Sell Active
type BSA struct {
	Bup  float64 `json:"bup"`
	Sdp  float64 `json:"sdp"`
	Bsr  float64 `json:"bsr"`
	Time string  `json:"t"`
}

// BSADay is Buy Sell Active by day
type BSADay struct {
	Ticker string  `json:"ticker"`
	Bsr    float64 `json:"bsr"`
	Date   string  `json:"d"`
}

// BSABind is binding for BSA
type BSABind struct {
	Ticker string `json:"ticker"`
	Data   []*BSA `json:"data"`
	Date   string `json:"d"`
}
