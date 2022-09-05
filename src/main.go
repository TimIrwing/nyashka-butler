package main

import (
	"github.com/TimIrwing/nyashka-butler/src/handlers"
	"github.com/TimIrwing/nyashka-butler/src/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(getToken())
	if err != nil {
		log.Panic(err)
	}

	m := messages.Messages{}
	m.ReadMessageFile()
	data := &handlers.Data{Messages: m}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	handlers.HandleUpdates(updates, bot, data)
}

func getToken() string {
	t, err := os.ReadFile("token")
	if err != nil {
		log.Panicf("Error getting token from file: %s\n", err)
	}
	return string(t)
}
