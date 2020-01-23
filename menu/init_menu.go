package menu

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"libretaxi/objects"
	"libretaxi/context"
	"log"
)

type InitMenuHandler struct {
}

func (handler *InitMenuHandler) Handle(user *objects.User, context *context.Context, message string) (callNext bool) {
	log.Println("Init menu")

	// Send welcome message
	msg := tgbotapi.NewMessage(user.UserId, "Welcome to LibreTaxi 2.0")
	context.Bot.Send(msg)

	// reset props and save to db
	user.MenuId = 0
	context.Repo.SaveUser(user)

	return true
}

func NewInitMenu() *InitMenuHandler {
	handler := &InitMenuHandler{}
	return handler
}
