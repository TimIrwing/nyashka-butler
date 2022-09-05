package commands

import (
	"github.com/TimIrwing/nyashka-butler/src/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetMessage(key, value string, m *tgbotapi.Message, messages messages.Messages) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(m.Chat.ID, "")
	messages[key] = value
	ok := messages.WriteMessageFile()
	if ok {
		msg.Text = "Список сообщений обновлён!"
	} else {
		msg.Text = "Не удалось обновить список сообщений, время смотреть логи"
	}
	msg.ReplyToMessageID = m.MessageID
	return msg
}
