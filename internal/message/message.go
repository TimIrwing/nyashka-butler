package message

import (
	"github.com/TimIrwing/nyashka-butler/internal/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

var EntityType = map[string]string{
	"mention":       "mention",       // (@username),
	"text_mention":  "text_mention",  // (for users without usernames)
	"hashtag":       "hashtag",       // (#hashtag),
	"cashtag":       "cashtag",       // ($USD),
	"command":       "bot_command",   // (/start@jobs_bot),
	"url":           "url",           // (https://telegram.org),
	"email":         "email",         // (do-not-reply@telegram.org),
	"number":        "phone_number",  // (+1-212-555-0123),
	"bold":          "bold",          // (bold text),
	"italic":        "italic",        // (italic text),
	"underline":     "underline",     // (underlined text),
	"strikethrough": "strikethrough", // (strikethrough text),
	"code":          "code",          // (monowidth string),
	"pre":           "pre",           // (monowidth block),
	"link":          "text_link",     // (for clickable text URLs),
}

type Bot interface {
	Send(c *Message)
}

type Message struct {
	ID      int
	Text    string
	ReplyTo int
	ChatID  int64
	Cmd     *tgbotapi.MessageEntity
	Bot     Bot
}

func (m *Message) parseEntities(list []tgbotapi.MessageEntity) {
	for _, e := range list {
		if e.Type == EntityType["command"] {
			m.Cmd = &e
		}
	}
}

func (m *Message) handleCommand() {
	if m.Cmd == nil {
		return
	}
	cmd := m.Text[m.Cmd.Offset+1 : m.Cmd.Length]
	cmd = strings.ToLower(cmd)
	opts := commands.CommandOptions{Message: m}

	trimmed := strings.Trim(m.Text[m.Cmd.Length:], " \n")
	for _, s := range strings.Split(trimmed, " ") {
		if len(s) > 0 {
			opts.Args = append(opts.Args, s)
		}
	}

	commands.Trigger(cmd, opts)
}

func From(m *tgbotapi.Message, bot Bot) Message {
	res := Message{
		Text:   m.Text,
		ID:     m.MessageID,
		ChatID: m.Chat.ID,
		Bot:    bot,
	}
	res.parseEntities(m.Entities)
	return res
}

func (m *Message) New(text string) commands.Message {
	return &Message{
		Text:   text,
		ChatID: m.ChatID,
		Bot:    m.Bot,
	}
}

func (m *Message) Handle() {
	m.handleCommand()
}

func (m *Message) Send() {
	m.Bot.Send(m)
}

func (m *Message) SetText(text string) {
	m.Text = text
}

func (m *Message) SetReply(id int) {
	m.ReplyTo = id
}

func (m *Message) GetID() int {
	return m.ID
}
