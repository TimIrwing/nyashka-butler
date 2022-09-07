package bot

import (
	"github.com/TimIrwing/nyashka-butler/internal/global"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	*tgbotapi.BotAPI
}

func New(token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panicf("Can't get botapi from telegram: %s", err)
	}
	return &Bot{bot}
}

func (bot *Bot) HandleUpdates(g *global.Global) {
	log.Printf("Authorized on account %s", bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	ch := bot.GetUpdatesChan(updateConfig)

	for {
		_u := <-ch
		u := Update{_u: &_u, bot: bot}
		u.Log()

		switch {
		case u.IsCommand():
			u.HandleCommand(g)

			// TODO case u.IsJoin():
			// TODO case u.IsLeave():
		}
	}
}
