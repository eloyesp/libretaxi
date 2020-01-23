package menu

import (
	"libretaxi/context"
	"libretaxi/objects"
	"log"
)

type Handler interface {
	Handle(user *objects.User, context *context.Context, message string) (success bool)
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

	//handler = NewInitMenu()
	handler = NewAskLocationMenu()
	handler.Handle(user, context, message)

	// call next on success
}