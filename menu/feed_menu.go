package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"libretaxi/objects"
	"libretaxi/context"
	"log"
)

type FeedMenuHandler struct {
}

func (handler *FeedMenuHandler) Handle(user *objects.User, context *context.Context, message *tgbotapi.Message) {
	log.Println("Feed menu")

	var buttons = []tgbotapi.KeyboardButton{
		{Text: "🔍 Search"},
		{Text: "🌎 Set location"},
	}
	msg := tgbotapi.NewMessage(user.UserId, "Done! You'll see new posts here. Use 🔍 to search for a 🚗 driver or 🤵 passenger.")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
	context.Bot.Send(msg)

	//user.MenuId = objects.Menu_AskLocation
	//context.Repo.SaveUser(user)
}

func NewFeedMenu() *FeedMenuHandler {
	handler := &FeedMenuHandler{}
	return handler
}
