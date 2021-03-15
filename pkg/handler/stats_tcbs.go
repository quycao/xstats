package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/quycao/xstats/pkg/model"
)

// StatsTCBS get transaction data of ticker
func StatsTCBS(ticker string) (*model.StatsResultTCBS, error) {
	url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/his", ticker)
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

	translogBind := model.TranslogTCBSBind{}
	err = json.Unmarshal(body, &translogBind)
	if err != nil {
		return nil, err
	}

	translogs := translogBind.Data

	sort.SliceStable(translogs, func(i, j int) bool {
		return translogs[i].Vol > translogs[j].Vol
	})

	result := &model.StatsResultTCBS{}
	if len(translogs) >= 50 {
		translogs = translogs[:50]
		var buyVol, selVol int64

		for _, val := range translogs {
			if err == nil {
				if val.Action == "BU" {
					buyVol = buyVol + val.Vol
				} else if val.Action == "SD" {
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

		result = &model.StatsResultTCBS{
			Ticker:     ticker,
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
