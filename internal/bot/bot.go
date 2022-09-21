package bot

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard"
	"github.com/TimIrwing/nyashka-butler/internal/bot/message"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type wrapper struct {
	api *tgbotapi.BotAPI
}

func Start(token string) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Can't get botapi from telegram: %s", err)
	}
	bot := wrapper{api: api}
	bot.start()
}

func (bot wrapper) start() {
	log.Printf("Authorized on account %s", bot.api.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	ch := bot.api.GetUpdatesChan(updateConfig)

	for {
		u := <-ch
		go bot.handleUpdate(u)
	}
}

func (bot wrapper) handleUpdate(u tgbotapi.Update) {
	switch {
	case u.Message != nil:
		m := message.From(u.Message)
		resp := m.Handle()
		bot.send(resp)
	}
}

func (bot wrapper) send(r *types.Response) {
	if r == nil {
		return
	}

	if m := r.Message; m != nil {
		res := tgbotapi.NewMessage(m.GetChatID(), m.GetText())
		res.ReplyToMessageID = m.GetReplyID()
		res.ReplyMarkup = keyboard.GetKeyboard(m.GetKeyboardPage())
		_, err := bot.api.Send(res)
		if err != nil {
			log.Printf("Couldn't send message: %s", err)
		}
	}
}
