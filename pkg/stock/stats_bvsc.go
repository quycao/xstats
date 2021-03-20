package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// StatsBVSC get transaction data of symbol
func StatsBVSC(symbol string) (*StatsResultBVSC, error) {
	url := fmt.Sprintf("https://online.bvsc.com.vn/datafeed/translogsnaps/%s", symbol)
	httpClient := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	translogBind := TranslogBVSCBind{}
	err = json.Unmarshal(body, &translogBind)
	if err != nil {
		return nil, err
	}

	translogs := translogBind.Data
	for idx, val := range translogs {
		vol, _ := strconv.Atoi(val.VolStr)
		translogs[idx].Vol = int64(vol)
	}

	sort.SliceStable(translogs, func(i, j int) bool {
		return translogs[i].Vol > translogs[j].Vol
	})

	result := &StatsResultBVSC{}
	if len(translogs) >= 0 {
		fivePct := len(translogs) * 5 / 100
		translogs = translogs[:fivePct]
		var buyVol, selVol int64

		for _, val := range translogs {
			if err == nil {
				if val.Type == "B" {
					buyVol = buyVol + val.Vol
				} else if val.Type == "S" {
					selVol = selVol + val.Vol
				}
			}
		}

		status := "Bình thường"
		suggestion := "Không"
		if float64(buyVol) >= float64(selVol)*1.5 {
			status = "Tích luỹ"
			suggestion = "Mua"
		} else if float64(selVol) >= float64(buyVol)*1.5 {
			status = "Phân phối"
			suggestion = "Bán"
		}

		result = &StatsResultBVSC{
			Symbol:     symbol,
			BuyVol:     buyVol,
			SellVol:    selVol,
			Status:     status,
			Suggestion: suggestion,
		}

		return result, nil

		// fmt.Printf("Buy: %d, Sell: %d", buyVol, selVol)
	} else {
		return nil, errors.New("There are less than 50 translog records")
		// fmt.Println("There are less than 50 translog records")
	}
}
