package handlers

import (
	"github.com/TimIrwing/nyashka-butler/src/commands"
	"github.com/TimIrwing/nyashka-butler/src/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Data struct {
	messages.Messages
}

func HandleUpdates(ch tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, data *Data) {
	for {
		u := <-ch
		if m := u.Message; m != nil {
			log.Printf("Message [%s] %s", m.From.UserName, m.Text)
			handleCommand(m, bot, data)
			handleJoinLeave(m, bot, data)
			handleReplyToMe(m, bot, data)
		}
	}
}

func handleCommand(m *tgbotapi.Message, bot *tgbotapi.BotAPI, data *Data) {
	command := ""
	finalText := ""
	for _, e := range m.Entities {
		if e.Type == "bot_command" {
			command = m.Text[:e.Length]
			finalText = strings.Trim(m.Text[e.Length:], " ")
			break
		}
	}
	if len(command) == 0 {
		return
	}
	switch command[1:] {
	case "welcome":
		msg := commands.SetMessage("welcome", finalText, m, data.Messages)
		bot.Send(msg)
	case "goodbye":
		msg := commands.SetMessage("goodbye", finalText, m, data.Messages)
		bot.Send(msg)
	case "reply":
		msg := commands.SetMessage("reply", finalText, m, data.Messages)
		bot.Send(msg)
	}
}

func handleJoinLeave(m *tgbotapi.Message, bot *tgbotapi.BotAPI, data *Data) {
	welcomeMsg := data.Messages["welcome"]
	if len(m.NewChatMembers) > 0 && len(welcomeMsg) > 0 {
		msg := tgbotapi.NewMessage(m.Chat.ID, welcomeMsg)
		msg.ReplyToMessageID = m.MessageID
		bot.Send(msg)
	}

	goodbyeMsg := data.Messages["goodbye"]
	if m.LeftChatMember != nil && len(goodbyeMsg) > 0 {
		msg := tgbotapi.NewMessage(m.Chat.ID, goodbyeMsg)
		msg.ReplyToMessageID = m.MessageID
		bot.Send(msg)
	}
}

func handleReplyToMe(m *tgbotapi.Message, bot *tgbotapi.BotAPI, data *Data) {
	replyMsg := data.Messages["reply"]
	if len(replyMsg) == 0 {
		return
	}
	reply := m.ReplyToMessage
	if reply == nil {
		return
	}

	if reply.From.ID == bot.Self.ID {
		msg := tgbotapi.NewMessage(m.Chat.ID, replyMsg)
		msg.ReplyToMessageID = m.MessageID
		bot.Send(msg)
	}
}
