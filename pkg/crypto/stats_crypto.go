package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// StatsCrypto2 stats symbol
// func StatsCrypto2(symbol string) (*StatsResultCrypto, error) {
// 	public := "jXAPW3oerzFmRez3VkY03OQDEexzqp8NPuLBQPYkQL7PZJBYNNiTIkj9WXH7pATt"
// 	secret := "E16YpdnRCgmO4ljqwrlMIiVhJSeQv6ffqZQimLtdfApHJTBW5elpBvIEEtAggYu6"

// 	client := binance.NewClient(public, secret)

// 	trades, err := client.NewHistoricalTradesService().Symbol(symbol).Limit(1000).Do(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	sort.SliceStable(trades, func(i, j int) bool {
// 		return trades[i].Quantity > trades[j].Quantity
// 	})

// 	result := &StatsResultCrypto{}
// 	if len(trades) > 0 {
// 		fivePct := int64(math.Round(float64(len(trades)) * 5 / 100))
// 		trades = trades[:fivePct]
// 		var buyVol, selVol, totalValue float64

// 		for _, val := range trades {
// 			if err == nil {
// 				if val.IsBuyerMaker == false {
// 					qty, _ := strconv.ParseFloat(val.Quantity, 64)
// 					price, _ := strconv.ParseFloat(val.Price, 64)
// 					buyVol = buyVol + qty
// 					totalValue = totalValue + qty*price
// 				} else if val.IsBuyerMaker == true {
// 					qty, _ := strconv.ParseFloat(val.Quantity, 64)
// 					price, _ := strconv.ParseFloat(val.Price, 64)
// 					selVol = selVol + qty
// 					totalValue = totalValue + qty*price
// 				}
// 			}
// 		}

// 		status := "Bình thường"
// 		suggestion := "Không"
// 		var buySellPct int64
// 		if selVol == 0 {
// 			selVol = 1
// 		}
// 		if buyVol == 0 {
// 			buyVol = 1
// 		}

// 		if selVol != 0 && buyVol > selVol {
// 			buySellPct = int64((buyVol - selVol) * 100 / selVol)
// 			if buySellPct > 100 {
// 				status = "Tích luỹ"
// 				suggestion = "Mua"
// 			}
// 		} else if buyVol != 0 && selVol > buyVol {
// 			buySellPct = int64((selVol - buyVol) * (-100) / buyVol)
// 			if buySellPct < -100 {
// 				status = "Phân phối"
// 				suggestion = "Bán"
// 			}
// 		}

// 		avgPrice := totalValue / (buyVol + selVol)

// 		result = &StatsResultCrypto{
// 			Symbol:     symbol,
// 			AvgPrice:   avgPrice,
// 			BuyVol:     buyVol,
// 			SellVol:    selVol,
// 			BuySellPct: buySellPct,
// 			Status:     status,
// 			Suggestion: suggestion,
// 		}

// 		return result, nil

// 		// fmt.Printf("Buy: %d, Sell: %d", buyVol, selVol)
// 	} else {
// 		return nil, errors.New("There are no translog records of " + symbol)
// 		// fmt.Println("There are less than 50 translog records")
// 	}
// }

// StatsCrypto stats symbol
func StatsCrypto(symbol string) (*StatsResultCrypto, error) {
	// public := "jXAPW3oerzFmRez3VkY03OQDEexzqp8NPuLBQPYkQL7PZJBYNNiTIkj9WXH7pATt"
	// secret := "E16YpdnRCgmO4ljqwrlMIiVhJSeQv6ffqZQimLtdfApHJTBW5elpBvIEEtAggYu6"
	// queryStr := fmt.Sprintf("symbol=%s", symbol)

	// // Create a new HMAC by defining the hash type and the key (as byte array)
	// h := hmac.New(sha256.New, []byte(secret))
	// // Write Data to it
	// h.Write([]byte(queryStr))
	// // Get result and encode as hexadecimal string
	// sha := hex.EncodeToString(h.Sum(nil))
	// fmt.Sprintf(sha)

	// url := fmt.Sprintf("https://api.binance.com/api/v3/trades?symbol=%s&signature=%s", symbol, sha)

	url := fmt.Sprintf("https://api.binance.com/api/v3/trades?symbol=%s&limit=1000", symbol)
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

	var trades []Trade
	err = json.Unmarshal(body, &trades)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(trades, func(i, j int) bool {
		return trades[i].Qty > trades[j].Qty
	})

	result := &StatsResultCrypto{}
	if len(trades) > 0 {
		fivePct := int64(math.Round(float64(len(trades)) * 5 / 100))
		trades = trades[:fivePct]
		var buyVol, selVol, totalValue float64

		for _, val := range trades {
			if err == nil {
				if val.IsBuyerMaker == false {
					qty, _ := strconv.ParseFloat(val.Qty, 64)
					price, _ := strconv.ParseFloat(val.Price, 64)
					buyVol = buyVol + qty
					totalValue = totalValue + qty*price
				} else if val.IsBuyerMaker == true {
					qty, _ := strconv.ParseFloat(val.Qty, 64)
					price, _ := strconv.ParseFloat(val.Price, 64)
					selVol = selVol + qty
					totalValue = totalValue + qty*price
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
			if buySellPct >= 100 {
				status = "Tích luỹ"
				suggestion = "Mua"
			}
		} else if buyVol != 0 && selVol > buyVol {
			buySellPct = int64((selVol - buyVol) * (-100) / buyVol)
			if buySellPct <= -100 {
				status = "Phân phối"
				suggestion = "Bán"
			}
		}

		avgPrice := totalValue / (buyVol + selVol)

		result = &StatsResultCrypto{
			Time:       time.Now(),
			Symbol:     symbol,
			AvgPrice:   avgPrice,
			BuyVol:     buyVol,
			SellVol:    selVol,
			BuySellPct: buySellPct,
			Status:     status,
			Suggestion: suggestion,
		}

		return result, nil

		// fmt.Printf("Buy: %d, Sell: %d", buyVol, selVol)
	} else {
		return nil, errors.New("There are no translog records of " + symbol)
		// fmt.Println("There are less than 50 translog records")
	}
}
