package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	kz "github.com/wesovilabs/koazee"
)

func PriceVolumeStats(symbol string, interval string, periodsBefore int) (*PriceVolumeStatsResult, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/klines?symbol=%s&interval=%s", symbol, interval)
	httpClient := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// req.Header.Add("X-MBX-APIKEY", public)
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

	var dataArr [][12]interface{}
	err = json.Unmarshal(body, &dataArr)
	if err != nil {
		return nil, err
	}

	if len(dataArr) > 32-periodsBefore {
		var pvData []*PriceVolume
		for i := len(dataArr) + periodsBefore - 31; i < len(dataArr)+periodsBefore; i++ {
			data := dataArr[i]
			priceVolume := &PriceVolume{}
			for j := 0; j < 12; j++ {
				switch j {
				case 0:
					priceVolume.OpenTime = time.Unix(int64(data[j].(float64)/1000), 0)
				case 1:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.OpenPrice = f
				case 2:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.HighPrice = f
				case 3:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.LowPrice = f
				case 4:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.ClosePrice = f
				case 5:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.Volume = f
				case 6:
					priceVolume.CloseTime = time.Unix(int64(data[j].(float64)), 0)
				case 7:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.QuoteAssetVolume = f
				case 8:
					priceVolume.NumberOfTrades = int64(data[j].(float64))
				case 9:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.TakerBuyBase = f
				case 10:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.TakerBuyQuote = f
				case 11:
					f, _ := strconv.ParseFloat(data[j].(string), 64)
					priceVolume.Ignore = f
				}
			}

			pvData = append(pvData, priceVolume)
		}

		// Update lastPV after change pvData
		lastPV := pvData[len(pvData)-1]
		secondLastPV := pvData[len(pvData)-2]

		// Get data of last 30 days before last date
		pvData = pvData[len(pvData)-31 : len(pvData)-1]

		pvs := kz.StreamOf(pvData)

		// Get Volumne of 10 last day
		tenLastPV := pvs.Take(20, 29).Do()
		avgVolume := tenLastPV.Reduce(func(acc float64, pv *PriceVolume) float64 {
			return acc + pv.Volume
		}).Float64()
		avgVolume = avgVolume / 10

		// Get max price within last 20 days (30 days on calendar)
		thirtyLastPV := pvs.Take(0, 29).Do()
		maxPrice := thirtyLastPV.Reduce(func(acc float64, pv *PriceVolume) (float64, error) {
			if acc < pv.ClosePrice {
				acc = pv.ClosePrice
			}
			return acc, nil
		}).Float64()
		avgVolumeChange := (lastPV.Volume - avgVolume) / avgVolume
		// avgVolumeChange = math.Round(avgVolumeChange*10000)/10000
		maxPriceChange := (lastPV.ClosePrice - maxPrice) / maxPrice
		// maxPriceChange = math.Round(maxPriceChange*10000)/10000
		priceChange := (lastPV.ClosePrice - secondLastPV.ClosePrice) / secondLastPV.ClosePrice
		// priceChange = math.Round(priceChange*10000)/10000

		result := &PriceVolumeStatsResult{
			Symbol:                    symbol,
			Period:                    interval,
			Price:                     lastPV.ClosePrice,
			Volume:                    lastPV.Volume,
			AvgVolume10Periods:        avgVolume,
			HighestPrice30Periods:     maxPrice,
			RatioChangeVol10Periods:   avgVolumeChange,
			RatioChangePrice30Periods: maxPriceChange,
			RatioChangePrice:          priceChange,
			Time:                      lastPV.OpenTime,
			Suggestion:                "None",
		}

		var priceChangeRatio, maxPriceChangeRatio float64
		if interval == "1d" {
			priceChangeRatio = 0.07
			maxPriceChange = 0.1
		} else if interval == "4h" {
			priceChangeRatio = 0.05
			maxPriceChange = 0.07
		} else if interval == "1h" {
			priceChangeRatio = 0.03
			maxPriceChange = 0.05
		} else if interval == "15m" {
			priceChangeRatio = 0.02
			maxPriceChange = 0.03
		}

		// Buy signal
		if (result.RatioChangePrice <= -1*priceChangeRatio || maxPriceChange <= -1*maxPriceChangeRatio) && lastPV.Volume >= avgVolume {
			result.Suggestion = "Buy"
		}

		// Sell signal
		if result.RatioChangePrice >= priceChangeRatio && float64(lastPV.Volume) >= float64(avgVolume)*2 {
			result.Suggestion = "Sell"
		}

		return result, nil
	} else {
		return nil, errors.New("There are not enough translog records of " + symbol)
	}
}
