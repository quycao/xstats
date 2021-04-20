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
	from := time.Now().AddDate(0, 0, daysBefore-180)
	to := time.Now().AddDate(0, 0, 1)
	url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/stock/bars-long-term?ticker=%s&type=stock&resolution=D&from=%d&to=%d", ticker, from.Unix(), to.Unix())
	// url := fmt.Sprintf("https://apiazure.tcbs.com.vn/public/stock-insight/v1/intraday/%s/pv?resolution=1440", ticker)
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
		// Get data of n days before if daysBefore is specify
		pvData = pvData[:len(pvData)+daysBefore]

		// Update lastPV after change pvData
		lastPV := pvData[len(pvData)-1]
		secondLastPV := pvData[len(pvData)-2]

		// Get data of last 30 days before last date
		pvData = pvData[len(pvData)-31 : len(pvData)-1]

		pvs := kz.StreamOf(pvData)

		if lastPV.Volume > 99999 {
			ratioChangePrice := (lastPV.Close - secondLastPV.Close) / secondLastPV.Close
			ratioChangePrice = math.Round(ratioChangePrice*10000) / 10000

			// Get Volumne of 10 last day
			tenLastPV := pvs.Take(20, 29).Do()
			avgVolume10Days := tenLastPV.Reduce(func(acc int64, pv *PriceVolume) (int64, error) {
				return acc + pv.Volume, nil
			}).Int64()
			avgVolume10Days = avgVolume10Days / 10

			// Get max price within last 20 days (30 days on calendar)
			thirtyLastPV := pvs.Take(0, 29).Do()
			maxPrice30Days := thirtyLastPV.Reduce(func(acc float64, pv *PriceVolume) (float64, error) {
				if acc < pv.Close {
					acc = pv.Close
				}
				return acc, nil
			}).Float64()
			volumeChange10Days := float64(lastPV.Volume-avgVolume10Days) / float64(avgVolume10Days)
			volumeChange10Days = math.Round(volumeChange10Days*10000) / 10000
			priceChange30Days := float64(lastPV.Close-maxPrice30Days) / float64(maxPrice30Days)
			priceChange30Days = math.Round(priceChange30Days*10000) / 10000

			maxPriceTenDays := tenLastPV.Reduce(func(acc float64, pv *PriceVolume) (float64, error) {
				if acc < pv.Close {
					acc = pv.Close
				}
				return acc, nil
			}).Float64()
			tren10Days := "Sideway"
			priceChangeTenDays := float64(lastPV.Close-maxPriceTenDays) / float64(maxPriceTenDays)
			if priceChangeTenDays > 0.027 {
				tren10Days = "Up"
			} else if priceChangeTenDays < -0.06 {
				tren10Days = "Down"
			}

			result := &PriceVolumeStatsResult{
				Ticker:                 ticker,
				Price:                  lastPV.Close,
				Volume:                 lastPV.Volume,
				Trend10Days:            tren10Days,
				AvgVolume10Days:        avgVolume10Days,
				HighestPrice30Days:     maxPrice30Days,
				RatioChangeVol10Days:   volumeChange10Days,
				RatioChangePrice30Days: priceChange30Days,
				RatioChangePrice:       ratioChangePrice,
				Date:                   strings.Split(lastPV.TradingDate, "T")[0],
				Suggestion:             "None",
			}

			// Buy signal
			if (tren10Days == "Down" || priceChange30Days <= -0.09) && float64(lastPV.Volume) >= float64(avgVolume10Days)*1.2 {
				result.Suggestion = "Buy"
			} else if (tren10Days == "Sideway" && float64(lastPV.Volume) >= float64(avgVolume10Days)*1.5) || (ratioChangePrice <= -0.025 && lastPV.Volume >= avgVolume10Days) {
				result.Suggestion = "Buy"
			} else
			// Sell signal
			if tren10Days == "Up" && ratioChangePrice >= 0.025 && float64(lastPV.Volume) >= float64(avgVolume10Days)*1.5 {
				result.Suggestion = "Sell"
			}

			return result, nil
		} else {
			result := &PriceVolumeStatsResult{
				Ticker:     ticker,
				Price:      lastPV.Close,
				Volume:     lastPV.Volume,
				Date:       strings.Split(lastPV.TradingDate, "T")[0],
				Suggestion: "None - Volume too small",
			}
			return result, nil
		}
	} else {
		return nil, errors.New("There are not enough translog records of " + ticker)
	}
}
