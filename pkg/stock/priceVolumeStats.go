package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// PriceVolumeStats get price volume data of ticker
func PriceVolumeStats(ticker string) (*string, error) {
	url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/pv?resolution=1440", ticker)
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

	pvBind := PriceVolumeBind{}
	err = json.Unmarshal(body, &pvBind)
	if err != nil {
		return nil, err
	}

	pvs := pvBind.Data

	if len(pvs) > 0 {
		// Get Price of last day
		lastPV := pvs[len(pvs)-1]

		if lastPV.Volume > 99999 {
			// Get Volumne of 5 last day
			fiveLastPV := pvs[len(pvs) - 6:len(pvs) - 1]
			var avgVolume int64 = 0
			for _, pv := range(fiveLastPV) {
				avgVolume = avgVolume + pv.Volume
			}
			avgVolume = avgVolume/5
	
			// Get max price within 20 days (30 days on calendar)
			startIdx := 0
			if startIdx > 21 {
				startIdx = len(pvs) - 21
			}
			thirtyLastPV := pvs[startIdx:len(pvs) - 1]
			var maxPrice int64 = 0;
			for _, pv := range(thirtyLastPV) {
				if pv.Price > maxPrice {
					maxPrice = pv.Price
				}
			}
			maxPriceChange := float64(lastPV.Price - maxPrice)/float64(maxPrice)
	
			// Buy signal
			var result string
			if (lastPV.RatioChangePrice < -0.03 || maxPriceChange < -0.9) && lastPV.Volume > avgVolume {
				result = "Buy"
			}
	
			// Sell signal
			if lastPV.ChangePrice > 0 && float64(lastPV.Volume) > float64(avgVolume)*1.5 {
				result = "Sell"
			}
			
			return &result, nil
		}
		return nil, nil
	} else {
		return nil, errors.New("There are no translog records of " + ticker)
	}
}
