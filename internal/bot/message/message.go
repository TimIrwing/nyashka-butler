package message

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard/pages"
	"github.com/TimIrwing/nyashka-butler/internal/bot/message/commands"
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
	"github.com/TimIrwing/nyashka-butler/internal/types"
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

type Message struct {
	id           int
	text         string
	replyID      int
	chatID       int64
	keyboardPage pages.KeyboardPage
	cmdEntity    *tgbotapi.MessageEntity
}

func (m *Message) parseEntities(list []tgbotapi.MessageEntity) {
	for _, e := range list {
		if e.Type == EntityType["command"] {
			m.cmdEntity = &e
		}
	}
}

func (m *Message) handleCommand() *types.Response {
	if m.cmdEntity == nil {
		return nil
	}

	opts := commands.CommandOptions{Message: m}
	trimmed := strings.Trim(m.text[m.cmdEntity.Length:], " \n")
	normalized := strings.ReplaceAll(trimmed, "\n", " ")
	for _, s := range strings.Split(normalized, " ") {
		if len(s) > 0 {
			opts.Args = append(opts.Args, s)
		}
	}

	cmd := m.text[m.cmdEntity.Offset+1 : m.cmdEntity.Length]
	cmd = strings.ToLower(cmd)
	resp := commands.Trigger(cmd, opts)
	return resp
}

func From(m *tgbotapi.Message) Message {
	res := Message{
		text:   m.Text,
		id:     m.MessageID,
		chatID: m.Chat.ID,
	}
	res.parseEntities(m.Entities)
	return res
}

func (m *Message) Handle() *types.Response {
	resp := m.handleCommand()
	return resp
}

func (m *Message) New(text string) interfaces.Message {
	return &Message{
		text:   text,
		chatID: m.chatID,
	}
}

func (m *Message) GetID() int {
	return m.id
}
func (m *Message) GetReplyID() int {
	return m.replyID
}
func (m *Message) GetChatID() int64 {
	return m.chatID
}
func (m *Message) GetText() string {
	return m.text
}
func (m *Message) GetKeyboardPage() pages.KeyboardPage {
	return m.keyboardPage
}

func (m *Message) SetReplyID(id int) {
	m.replyID = id
}
func (m *Message) SetText(text string) {
	m.text = text
}
func (m *Message) SetKeyboardPage(p pages.KeyboardPage) {
	m.keyboardPage = p
}
