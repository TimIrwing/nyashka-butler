package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Update struct {
	_u  *tgbotapi.Update
	bot *Bot
}

type Response struct {
	text    string
	replyTo int
}

func (u *Update) Log() {
	m := u._u.Message
	if m == nil {
		return
	}
	log.Printf("Message [%s] %s", m.From.UserName, m.Text)
}

func (u *Update) Respond(r Response) {
	// TODO
}
