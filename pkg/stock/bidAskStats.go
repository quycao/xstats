package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BidAskStats get over by percent data of ticker
func BidAskStats(ticker string) (*BidAskPercentByDay, error) {
	url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/bid-ask?mode=baAll", ticker)
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

	bidAskBind := BidAskBind{}
	err = json.Unmarshal(body, &bidAskBind)
	if err != nil {
		return nil, err
	}

	bidAskLog := bidAskBind.OverBidAskLog

	if len(bidAskLog) > 0 {
		var sum float64
		for _, val := range bidAskLog {
			sum = sum + val.Aobp
		}
		avgBA := sum / float64(len(bidAskLog))
		baPercentByDay := &BidAskPercentByDay{
			Ticker:    ticker,
			OBPercent: avgBA,
			Date:      bidAskBind.Date,
		}
		return baPercentByDay, nil
	} else {
		return nil, errors.New("There are no translog records of " + ticker)
	}
}
