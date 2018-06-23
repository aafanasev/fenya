package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	// create a bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_FENYA_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// set webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://secret-shore-80825.herokuapp.com/" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	//
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	//
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed] %s", info.LastErrorMessage)
	}

	// set webhook listener
	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}

}
