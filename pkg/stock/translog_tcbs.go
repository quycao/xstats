package stock

// TranslogTCBS is data type of Transaction history in day (lịch sử khớp lệnh)
type TranslogTCBS struct {
	Price            int64   `json:"p"`
	Vol              int64   `json:"v"`
	ChangePrice      float64 `json:"cp"`
	RatioChangePrice float64 `json:"rcp"`
	Action           string  `json:"a"`
	Time             string  `json:"t"`
}

// TranslogTCBSBind is binding for translog
type TranslogTCBSBind struct {
	Ticker string          `json:"ticker"`
	Data   []*TranslogTCBS `json:"data"`
	Date   string          `json:"d"`
}
