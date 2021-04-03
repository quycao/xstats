package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// TranslogStats get transaction data of ticker
func TranslogStats(ticker string) (*TranslogDay, error) {
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

	translogBind := TranslogBind{}
	err = json.Unmarshal(body, &translogBind)
	if err != nil {
		return nil, err
	}

	translogs := translogBind.Data

	if len(translogs) > 0 {
		var buyVol, selVol, totalValue, totalChangeValue int64

		for _, val := range translogs {
			if err == nil {
				if val.Action == "BU" {
					buyVol = buyVol + val.Vol
					totalValue = totalValue + val.Vol*val.Price
					totalChangeValue = totalChangeValue + val.Vol*int64(val.ChangePrice)
				} else if val.Action == "SD" {
					selVol = selVol + val.Vol
					totalValue = totalValue + val.Vol*val.Price
					totalChangeValue = totalChangeValue + val.Vol*int64(val.ChangePrice)
				}
			}
		}

		translogDay := &TranslogDay{
			Ticker:         ticker,
			TotalVol:       (buyVol + selVol) / 2,
			AvgPrice:       totalValue / (buyVol + selVol),
			AvgChangePrice: float64(totalChangeValue / (buyVol + selVol)),
			Date:           translogBind.Date,
		}

		return translogDay, nil
	} else {
		return nil, errors.New("There are no translog records of " + ticker)
	}
}
