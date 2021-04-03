package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BSAStats get buy sell active data of ticker
func BSAStats(ticker string) (*BSADay, error) {
	url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/bsa", ticker)
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

	bsaBind := BSABind{}
	err = json.Unmarshal(body, &bsaBind)
	if err != nil {
		return nil, err
	}

	bsas := bsaBind.Data

	if len(bsas) > 0 {
		dayBSA := bsas[len(bsas)-1]
		bsaDay := &BSADay{
			Ticker: ticker,
			Bsr:    dayBSA.Bsr,
			Date:   bsaBind.Date,
		}
		return bsaDay, nil
	} else {
		return nil, errors.New("There are no translog records of " + ticker)
	}
}
