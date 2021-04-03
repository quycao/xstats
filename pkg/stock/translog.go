package stock

// Translog is data type of Transaction history in day (lịch sử khớp lệnh)
type Translog struct {
	Price            int64   `json:"p"`
	Vol              int64   `json:"v"`
	ChangePrice      float64 `json:"cp"`
	RatioChangePrice float64 `json:"rcp"`
	Action           string  `json:"a"`
	Time             string  `json:"t"`
}

// TranslogDay is summary translog of day
type TranslogDay struct {
	Ticker              string  `json:"ticker"`
	TotalVol            int64   `json:"v"`
	AvgPrice            int64   `json:"p"`
	AvgChangePrice      float64 `json:"cp"`
	AvgRatioChangePrice float64 `json:"rcp"`
	Date                string  `json:"d"`
}

// TranslogBind is binding for translog
type TranslogBind struct {
	Ticker string      `json:"ticker"`
	Data   []*Translog `json:"data"`
	Date   string      `json:"d"`
}
