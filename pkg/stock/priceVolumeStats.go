package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	kz "github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
)

// PriceVolumeStats get price volume data of ticker
func PriceVolumeStats(ticker string) (*string, error) {
	url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/pv?resolution=1440", ticker)
	// url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/pv?resolution=1440", "PVD")
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

	if len(pvBind.Data) > 32 {
		pvBind.Data = pvBind.Data[:len(pvBind.Data) - 1]
		itemCount := len(pvBind.Data)

		//init the loc
		loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		timeInLoc := time.Now().In(loc)
		var pvs stream.Stream
		if (timeInLoc.Hour() > 9 && timeInLoc.Hour() < 15) {
			pvs = kz.StreamOf(pvBind.Data[itemCount - 32:itemCount - 1])
		} else {
			pvs = kz.StreamOf(pvBind.Data[itemCount - 31:itemCount])
		}

		// Get Price of last day
		lastPV := pvs.Last().Val().(*PriceVolume)

		if lastPV.Volume > 99999 {
			// Get Volumne of 10 last day
			tenLastPV := pvs.Take(20, 29).Do()
			avgVolume := tenLastPV.Reduce(func (acc int64, pv *PriceVolume) int64 {
				return acc + pv.Volume
			}).Int64()
			avgVolume = avgVolume/10
	
			// Get max price within last 20 days (30 days on calendar)
			thirtyLastPV := pvs.Take(0, 29).Do()
			maxPrice := thirtyLastPV.Reduce(func (acc int64, pv *PriceVolume) (int64, error) {
				if acc < pv.Price {
					acc = pv.Price
				}
				return acc, nil
			}).Int64()
			maxPriceChange := float64(lastPV.Price - maxPrice)/float64(maxPrice)
	
			// Buy signal
			var result string
			if (lastPV.RatioChangePrice < -0.03 || maxPriceChange < -0.07) && lastPV.Volume > avgVolume {
				result = "Buy"
			}
	
			// Sell signal
			if lastPV.RatioChangePrice > 0.03 && float64(lastPV.Volume) > float64(avgVolume)*2 {
				result = "Sell"
			}
			
			return &result, nil
		}
		return nil, nil
	} else {
		return nil, errors.New("There are no translog records of " + ticker)
	}
}
