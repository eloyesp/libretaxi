package menu

import (
	"libretaxi/context"
	"libretaxi/objects"
	"log"
)

type MenuId int

const (
	Init MenuId = 0
	AskLocation MenuId = 100
	Feed MenuId = 200
	Post MenuId = 300
)

type Handler interface {
	Handle(user *objects.User, context *context.Context, message string) (callNext bool)
}

func HandleMessage(context *context.Context, userId int64, message string) {
	user := context.Repo.FindUser(userId)

	if user == nil {
		user = &objects.User{
			UserId: userId,
			MenuId: 0,
		}
	}

	//fmt.Printf("%+v\n", user)
	log.Printf("Message: '%s'", message)

	var handler Handler

	handler = NewInitMenu()
	handler.Handle(user, context, message)
}