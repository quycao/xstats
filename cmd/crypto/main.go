package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/quycao/xstats/pkg/crypto"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// fmt.Print("Input ticker symbol: ")
	// reader := bufio.NewReader(os.Stdin)
	// // ReadString will block until the delimiter is entered
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("An error occured while reading input. Please try again", err)
	// 	return
	// }

	// // remove the delimeter from the string
	// input = strings.TrimSuffix(input, "\n")
	// input = strings.ToUpper(input)
	// input = fmt.Sprintf("%sBUSD", input)
	// s.Every(1).Minutes().Do(stats, input)
	// s.StartBlocking()

	// stats("BNB")

	port := os.Getenv("PORT")            // sets automatically
	publicURL := os.Getenv("PUBLIC_URL") // you must add it to your config vars
	token := os.Getenv("TOKEN")          // you must add it to your config vars

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	bot, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	var senders []*tb.User
	bot.Handle("/sub", func(m *tb.Message) {
		senders = append(senders, m.Sender)
		// bot.Send(m.Sender, fmt.Sprintf("Your chat id: %d", m.Chat.ID))
		bot.Send(m.Sender, fmt.Sprintf("Your have followed %s", m.Payload))

		s := gocron.NewScheduler(time.UTC)
		s.Every(1).Minutes().Tag(m.Payload).Do(statsAndSend, m.Payload, bot, m.Sender)
		// s.Every(1).Minutes().Do(statsAndSend, "ETHBUSD", bot, m.Sender)
		// s.Every(1).Minutes().Do(statsAndSend, "BNBBUSD", bot, m.Sender)
		// s.Every(1).Minutes().Do(statsAndSend, "KNCBUSD", bot, m.Sender)
		s.StartAsync()

		// statsResult, err := crypto.StatsCrypto("BNBBUSD")
		// if err != nil {
		// 	log.Println(err)
		// 	bot.Send(m.Sender, err)
		// } else {
		// 	log.Println(statsResult.ToString())
		// 	bot.Send(m.Sender, statsResult.ToString())
		// }
	})

	bot.Handle("/unsub", func(m *tb.Message) {
		s := gocron.NewScheduler(time.UTC)
		s.RemoveByTag(m.Payload)
		bot.Send(m.Sender, fmt.Sprintf("Your have unfollowed %s", m.Payload))
	})

	// for _, sender := range senders {
	// 	s.Every(1).Minutes().Do(statsAndSend, "BNBBUSD", bot, sender)
	// }
	// s.StartAsync()

	bot.Start()
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

func statsAndSend(symbol string, bot *tb.Bot, user *tb.User) {
	statsResult, err := crypto.StatsCrypto(symbol)
	if err != nil {
		bot.Send(user, err)
	} else {
		bot.Send(user, statsResult.ToString())
	}
}
