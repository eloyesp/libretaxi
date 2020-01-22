package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // important
	"libretaxi/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"libretaxi/repository"
	"log"
)

type Context struct {
	bot *tgbotapi.BotAPI
	repo *repository.Repository
}

func initContext() *Context {
	config.Init("libretaxi")
	log.Printf("Using '%s' telegram token...\n", config.C().Telegram_Token)
	log.Printf("Using '%s' database connection string...", config.C().Db_Conn_Str)

	context := &Context{}

	bot, err := tgbotapi.NewBotAPI(config.C().Telegram_Token)
	if err != nil {
		log.Panic(err)
	}
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	db, err := sql.Open("postgres", config.C().Db_Conn_Str)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Successfully connected to the database...")
	}
	db.Query("SELECT 1")

	context.bot = bot
	context.repo = repository.NewRepository(db)
	return context
}

func main() {
	context := initContext()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := context.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%d - %s] %s", update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

		user := context.repo.FindUser(update.Message.Chat.ID)
		fmt.Printf("%+v\n", user)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		context.bot.Send(msg)
	}
}
