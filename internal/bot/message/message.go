package message

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/message/commands"
	"github.com/TimIrwing/nyashka-butler/internal/bot/message/text-interaction"
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Message struct {
	id              int
	text            string
	user            types.User
	replyID         int
	replyUserID     int64
	chatID          int64
	cmd             string
	cmdArgs         []string
	keyboardPage    types.KeyboardPage
	textInteraction types.TextInteractionName
}

func (m *Message) parseCmdArgs(args string) {
	var res []string
	for _, cur := range strings.Split(args, " ") {
		if len(cur) > 0 {
			res = append(res, cur)
		}
	}
	m.cmdArgs = res
}

func (m *Message) parseReply(replyMsg *tgbotapi.Message) {
	if replyMsg == nil {
		return
	}
	m.replyID = replyMsg.MessageID
}

func (m *Message) New(text string) interfaces.Message {
	return &Message{
		text:   text,
		chatID: m.chatID,
	}
}

func From(m *tgbotapi.Message, botInfo *types.BotInfo) *Message {
	res := &Message{
		id:              m.MessageID,
		text:            m.Text,
		user:            types.User{ID: m.From.ID, UserName: m.From.UserName},
		chatID:          m.Chat.ID,
		cmd:             strings.ToLower(m.Command()),
		textInteraction: text_interaction.GetName(m, botInfo),
	}
	res.parseReply(m.ReplyToMessage)
	res.parseCmdArgs(m.CommandArguments())
	return res
}

func (m *Message) Log() {
	log.Printf("[Message] ID:%d chatID:%d | userID:%d | user:@%s\n%s", m.id, m.chatID, m.user.ID, m.user.UserName, m.text)
}

func (m *Message) Handle(s interfaces.Settings) interfaces.Message {
	var resp interfaces.Message

	switch {
	case len(m.cmd) > 0:
		resp = commands.Handle(m.cmd, commands.CommandOptions{
			Message: m,
			Args:    m.cmdArgs,
		})
	default:
		resp = text_interaction.Handle(m, s)
	}

	return resp
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
func (m *Message) GetKeyboardPage() types.KeyboardPage {
	return m.keyboardPage
}
func (m *Message) GetTextInteraction() types.TextInteractionName {
	return m.textInteraction
}

func (m *Message) SetReplyID(id int) {
	m.replyID = id
}
func (m *Message) SetText(text string) {
	m.text = text
}
func (m *Message) SetKeyboardPage(p types.KeyboardPage) {
	m.keyboardPage = p
}
