package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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
	if len(translogs) > 0 {
		fivePct := int64(math.Round(float64(len(translogs)) * 5 / 100))
		translogs = translogs[:fivePct]
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
		var buySellPct int64
		if selVol == 0 {
			selVol = 1
		}
		if buyVol == 0 {
			buyVol = 1
		}

		if selVol != 0 && buyVol > selVol {
			buySellPct = int64((buyVol - selVol) * 100 / selVol)
			if buySellPct > 50 {
				status = "Tích luỹ"
				suggestion = "Mua"
			}
		} else if buyVol != 0 && selVol > buyVol {
			buySellPct = int64((selVol - buyVol) * (-100) / buyVol)
			if buySellPct < -50 {
				status = "Phân phối"
				suggestion = "Bán"
			}
		}

		result = &model.StatsResultTCBS{
			Ticker:     ticker,
			BuyVol:     buyVol,
			SellVol:    selVol,
			BuySellPct: buySellPct,
			Status:     status,
			Suggestion: suggestion,
		}

		return result, nil

		// fmt.Printf("Buy: %d, Sell: %d", buyVol, selVol)
	} else {
		return nil, errors.New("There are no translog records of " + ticker)
		// fmt.Println("There are less than 50 translog records")
	}
}
