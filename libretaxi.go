package main

import (
	"fmt"
	"libretaxi/config"
)

func main() {
	config.Init("libretaxi")
	fmt.Printf("Using %s telegram token", config.C().Telegram_Token)
}
