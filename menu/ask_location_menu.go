package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"libretaxi/objects"
	"libretaxi/context"
	"log"
)

type AskLocationMenuHandler struct {
	user *objects.User
	context *context.Context
	message string
}

func (handler *AskLocationMenuHandler) saveLocation() (success bool) {
	//user.MenuId = 0
	//context.Repo.SaveUser(user)

	return false
}

func (handler *AskLocationMenuHandler) Handle(user *objects.User, context *context.Context, message string) (success bool) {
	log.Println("Ask location menu")

	handler.user = user
	handler.context = context
	handler.message = message

	user.MenuId = objects.Menu_AskLocation
	context.Repo.SaveUser(user)

	if len(message) > 0 && handler.saveLocation() {
		return true
	} else {
		var buttons = []tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButtonLocation("Next"),
		}

		msg := tgbotapi.NewMessage(user.UserId, "Click \"Next\" (from mobile phone) to share your location.")
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
		context.Bot.Send(msg)
	}

	return false
}

func NewAskLocationMenu() *AskLocationMenuHandler {
	handler := &AskLocationMenuHandler{}
	return handler
}
