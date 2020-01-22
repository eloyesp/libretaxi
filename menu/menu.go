package menu

import (
	"fmt"
	"libretaxi/context"
	"log"
)

func HandleMessage(context *context.Context, userId int64, message string) {
	user := context.Repo.FindUser(userId)
	fmt.Printf("%+v\n", user)
	log.Println(message)
}