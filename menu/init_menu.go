package menu

import (
	"libretaxi/objects"
	"libretaxi/context"
	"log"
)

type InitMenuHandler struct {
}

func (handler *InitMenuHandler) Handle(user *objects.User, context *context.Context, message string) (callNext bool) {
	log.Println("init menu")
	return false
}

func NewInitMenu() *InitMenuHandler {
	handler := &InitMenuHandler{}
	return handler
}
