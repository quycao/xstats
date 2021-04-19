package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"

	kz "github.com/wesovilabs/koazee"
)

// PriceVolumeStats get price volume data of ticker
func PriceVolumeStats(ticker string, daysBefore int) (*PriceVolumeStatsResult, error) {
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

	if len(pvBind.Data) > 32-daysBefore {
		pvData := pvBind.Data

		// Get Price of last day
		lastPV := pvData[len(pvData)-1]
		secondLastPV := pvData[len(pvData)-2]

		if strings.Split(lastPV.Date, " ")[0] == strings.Split(secondLastPV.Date, " ")[0] {
			pvData = pvData[:len(pvData)-1]
		}

		// Get data of n days before if daysBefore is specify
		pvData = pvData[:len(pvData)+daysBefore]

		// Update lastPV after change pvData
		lastPV = pvData[len(pvData)-1]
		secondLastPV = pvData[len(pvData)-2]

		// Get data of last 30 days before last date
		pvData = pvData[len(pvData)-31 : len(pvData)-1]

		// //init the loc
		// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		// timeInLoc := time.Now().In(loc)

		// var pvs stream.Stream
		// if (timeInLoc.Hour() > 9 && timeInLoc.Hour() < 15) {
		// 	pvs = kz.StreamOf(pvData[len(pvData) - 32:len(pvData) - 1])
		// } else {
		// 	pvs = kz.StreamOf(pvData[len(pvData) - 31:len(pvData)])
		// }

		pvs := kz.StreamOf(pvData)

		if lastPV.Volume > 99999 {
			// Get Volumne of 10 last day
			tenLastPV := pvs.Take(20, 29).Do()
			avgVolume := tenLastPV.Reduce(func(acc int64, pv *PriceVolume) int64 {
				return acc + pv.Volume
			}).Int64()
			avgVolume = avgVolume / 10

			// Get max price within last 20 days (30 days on calendar)
			thirtyLastPV := pvs.Take(0, 29).Do()
			maxPrice := thirtyLastPV.Reduce(func(acc int64, pv *PriceVolume) (int64, error) {
				if acc < pv.Price {
					acc = pv.Price
				}
				return acc, nil
			}).Int64()
			avgVolumeChange := float64(lastPV.Volume-avgVolume) / float64(avgVolume)
			avgVolumeChange = math.Round(avgVolumeChange*10000) / 10000
			maxPriceChange := float64(lastPV.Price-maxPrice) / float64(maxPrice)
			maxPriceChange = math.Round(maxPriceChange*10000) / 10000

			direction := "sideway"
			firstTenPV := tenLastPV.First().Val().(*PriceVolume)
			if float64(secondLastPV.Price-firstTenPV.Price)/float64(firstTenPV.Price) >= 0.1 {
				direction = "up"
			} else if float64(secondLastPV.Price-firstTenPV.Price)/float64(firstTenPV.Price) <= -0.1 {
				direction = "down"
			}

			result := &PriceVolumeStatsResult{
				Ticker:                 ticker,
				Price:                  lastPV.Price,
				Volume:                 lastPV.Volume,
				AvgVolume10Days:        avgVolume,
				HighestPrice30Days:     maxPrice,
				RatioChangeVol10Days:   avgVolumeChange,
				RatioChangePrice30Days: maxPriceChange,
				RatioChangePrice:       lastPV.RatioChangePrice,
				Date:                   lastPV.Date,
				Suggestion:             "None",
			}

			// Buy signal
			if direction == "down" && float64(lastPV.Volume) >= float64(avgVolume)*1.5 {
				result.Suggestion = "Buy"
			} else if direction == "sideway" && (lastPV.RatioChangePrice <= -0.027 || maxPriceChange <= -0.07) && lastPV.Volume >= avgVolume {
				result.Suggestion = "Buy"
			}

			// Sell signal
			if direction == "up" && lastPV.RatioChangePrice >= 0.027 && float64(lastPV.Volume) >= float64(avgVolume)*1.5 {
				result.Suggestion = "Sell"
			}

			return result, nil
		} else {
			result := &PriceVolumeStatsResult{
				Ticker:     ticker,
				Price:      lastPV.Price,
				Volume:     lastPV.Volume,
				Date:       lastPV.Date,
				Suggestion: "None - Volume too small",
			}
			return result, nil
		}
	} else {
		return nil, errors.New("There are not enough translog records of " + ticker)
	}
}
