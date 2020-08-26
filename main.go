package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+event.Source.UserID)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func postToBack() {
	data := url.Values{
		"action":       {"CO2O"},
		"message_type": {"line"},
		"user_id":      {"5622512"},
		"tel":          {"123456"},
		"name":         {"gogo610"},
		"birthday":     {"1010101"},
	}
	resp, err := http.PostForm("http://192.168.100.48:8010/610_is_good/user_info", data)

	if err != nil {
		panic(err)
	}

	var res interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)
}
