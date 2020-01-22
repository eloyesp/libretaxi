package main

import (
	"libretaxi/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	config.Init("libretaxi")
	log.Printf("Using '%s' telegram token...\n", config.C().Telegram_Token)

	bot, err := tgbotapi.NewBotAPI(config.C().Telegram_Token)
	if err != nil {
		log.Panic(err)
	}
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%d - %s] %s", update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
