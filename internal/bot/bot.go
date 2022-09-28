package bot

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard"
	"github.com/TimIrwing/nyashka-butler/internal/bot/message"
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
	"github.com/TimIrwing/nyashka-butler/internal/settings"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type wrapper struct {
	api *tgbotapi.BotAPI
	db  *mongodb.DB
}

func Start(token string, db *mongodb.DB) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Can't get botapi from telegram: %s", err)
	}
	bot := wrapper{api: api, db: db}
	bot.start()
}

func (bot wrapper) start() {
	log.Printf("Authorized on account %s", bot.api.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.api.GetUpdatesChan(updateConfig)

	for u := range updates {
		go bot.handleUpdate(u)
	}
}

func (bot wrapper) handleUpdate(u tgbotapi.Update) {
	var resp []interfaces.Message

	switch {
	case u.Message != nil:
		s := settings.New(bot.db, u.Message.Chat.ID)
		resp = append(resp, message.From(u.Message, bot.GetInfo()).Handle(s))
	}

	for _, r := range resp {
		if r != nil {
			go bot.send(r)
		}
	}
}

func (bot wrapper) send(m interfaces.Message) {
	text := m.GetText()
	if len(text) == 0 {
		log.Printf("Bot is trying to send empty text")
		return
	}

	res := tgbotapi.NewMessage(m.GetChatID(), text)
	res.ReplyToMessageID = m.GetReplyID()
	res.ReplyMarkup = keyboard.GetKeyboard(m.GetKeyboardPage())
	_, err := bot.api.Send(res)
	if err != nil {
		log.Printf("Couldn't send message: %s", err)
	}
}

func (bot wrapper) GetInfo() *types.BotInfo {
	return &types.BotInfo{
		ID:       bot.api.Self.ID,
		UserName: bot.api.Self.UserName,
	}
}