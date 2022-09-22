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
	cmd          string
	cmdArgs      []string
}

func (m *Message) handleCommand() *types.Response {
	if len(m.cmd) == 0 {
		return nil
	}
	return commands.Trigger(m.cmd, commands.CommandOptions{
		Message: m,
		Args:    m.cmdArgs,
	})
}

func From(m *tgbotapi.Message) Message {
	var args []string
	for _, cur := range strings.Split(m.CommandArguments(), " ") {
		if len(cur) > 0 {
			args = append(args, cur)
		}
	}

	return Message{
		text:    m.Text,
		id:      m.MessageID,
		chatID:  m.Chat.ID,
		cmd:     strings.ToLower(m.Command()),
		cmdArgs: args,
	}
}

func (m *Message) Handle() *types.Response {
	return m.handleCommand()
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
