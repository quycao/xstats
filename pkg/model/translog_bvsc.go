package model

// TranslogBVSC is data type of Transaction history in day (lịch sử khớp lệnh)
type TranslogBVSC struct {
	Seq         int64  `json:"sequenceMsg"`
	Symbol      string `json:"symbol"`
	Date        string `json:"tradingdate"`
	Time        string `json:"formattedTime"`
	MatchPrice  string `json:"formattedMatchPrice"`
	ChangePrice string `json:"formattedChangeValue"`
	VolStr      string `json:"formattedVol"`
	TotalVol    string `json:"formattedAccVol"`
	TotalVal    string `json:"formattedAccVal"`
	Type        string `json:"lastColor"`
	Vol         int64  `json:"vol"`
}

// TranslogBVSC is binding for translog
type TranslogBVSCBind struct {
	Status string          `json:"s"`
	Data   []*TranslogBVSC `json:"d"`
}
