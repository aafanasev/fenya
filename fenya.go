package fenya

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

var bot *tgbotapi.BotAPI

func Init(token, baseUrl string, debug bool) error {
	var err error

	// create a bot
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	log.Printf("[%s] connected", bot.Self.UserName)

	// debug mode
	bot.Debug = debug

	// set webhook
	webhookUrl := baseUrl + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl))
	if err != nil {
		return err
	}

	log.Printf("[%s] set webhook %s", bot.Self.UserName, webhookUrl)

	return err
}

func Handle(update *tgbotapi.Update) {

	if update.Message.NewChatMembers != nil {
		// say hello
		const greetingStickerId = "CAADAgADQgADaJpdDJ7cv5VPztHoAg"
		msg := tgbotapi.NewStickerShare(update.Message.Chat.ID, greetingStickerId)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}

}
