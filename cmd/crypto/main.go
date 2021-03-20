package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/quycao/xstats/pkg/crypto"
)

func main() {
	fmt.Print("Input ticker symbol: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	input = strings.ToUpper(input)
	input = fmt.Sprintf("%sBUSD", input)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(stats, input)
	s.StartBlocking()

	// stats()
}

func stats(symbol string) {
	statsResult, err := crypto.StatsCrypto(symbol)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(statsResult.ToString())
	}

	// url := fmt.Sprintf("https://api.binance.com/api/v3/trades?symbol=BNBBUSD")
	// httpClient := http.Client{Timeout: time.Second * 5}
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// // req.Header.Add("X-MBX-APIKEY", public)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// res, err := httpClient.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// body, err := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body[:]))
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
