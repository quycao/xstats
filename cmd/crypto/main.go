package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/quycao/xstats/pkg/crypto"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main1() {
	result, err := crypto.PriceVolumeStats("BNBBUSD", "1h", -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.ToString())
	}
}

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

	// bot.Handle("/", func(m *tb.Message) {
	// 	bot.Send(m.Sender, "Hi!")
	// })

	bot.Handle("/start", func(m *tb.Message) {
		bot.Send(m.Sender, fmt.Sprint("Hi!\nUse '/hi symbol' to follow\nUse '/remove symbol' to unfollow\nUse '/list' to get followed symbols"))
	})

	s := gocron.NewScheduler(time.UTC)
	bot.Handle("/hi", func(m *tb.Message) {
		input := strings.ToUpper(m.Payload)
		symbols := strings.Split(input, ",")
		for _, symbol := range symbols {
			symbol = strings.TrimSpace(symbol)
			tag := fmt.Sprintf("%s: %s", m.Sender.FirstName, symbol)
			pair := fmt.Sprintf("%sBUSD", symbol)
			s.Every(5).Minutes().Tag(tag).Do(statsAndSend, pair, bot, m.Sender, true)
		}
		s.StartAsync()
		bot.Send(m.Sender, fmt.Sprintf("You have followed %s", input))

		// statsResult, err := crypto.StatsCrypto("BNBBUSD")
		// if err != nil {
		// 	log.Println(err)
		// 	bot.Send(m.Sender, err)
		// } else {
		// 	log.Println(statsResult.ToString())
		// 	bot.Send(m.Sender, statsResult.ToString())
		// }
	})

	bot.Handle("/remove", func(m *tb.Message) {
		bot.Send(m.Sender, fmt.Sprintf("You have unfollowed %s", m.Payload))
		// s.Stop()
		s.RemoveByTag(m.Payload)
		s.StartAsync()
	})

	bot.Handle("/list", func(m *tb.Message) {
		if len(s.Jobs()) > 0 {
			var jobs string
			for _, job := range s.Jobs() {
				jobs = jobs + strings.Join(job.Tags(), ", ") + "\n"
			}
			bot.Send(m.Sender, jobs)
		} else {
			bot.Send(m.Sender, "You don't follow any symbol")
		}
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text == "hi" {
			bot.Send(m.Sender, fmt.Sprint("Hi!\nInput 'symbol' to get data"))
		} else {
			input := strings.ToUpper(fmt.Sprintf("%sBUSD", m.Text))
			statsAndSend(input, bot, m.Sender, false)
		}
	})

	// for _, sender := range senders {
	// 	s.Every(1).Minutes().Do(statsAndSend, "BNBBUSD", bot, sender)
	// }
	// s.StartAsync()

	// Start a web server in other go routine
	// go func() {
	// 	mux := http.NewServeMux()
	// 	mux.HandleFunc("/home", home)
	// 	fmt.Printf("Starting server on :%s\n", port)
	// 	err := http.ListenAndServe(":"+port, mux)
	// 	fmt.Println(err)
	// }()

	bot.Start()
	fmt.Println("Bot started")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
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

func statsAndSend(symbol string, bot *tb.Bot, user *tb.User, isActionOnly bool) {
	// statsResult, err := crypto.StatsCrypto(symbol)
	statsResult, err := crypto.PriceVolumeStats(symbol, "1h", 0)
	if err != nil {
		log.Println(err)
		bot.Send(user, err)
	} else {
		log.Println(statsResult.ToString())
		if isActionOnly {
			if statsResult.Suggestion != "None" {
				bot.Send(user, statsResult.ToString())
			}
		} else {
			bot.Send(user, statsResult.ToString())
		}
	}
}
