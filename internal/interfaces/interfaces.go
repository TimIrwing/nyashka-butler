package interfaces

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard/pages"
)

type Message interface {
	New(string) Message
	GetID() int
	GetReplyID() int
	GetChatID() int64
	GetText() string
	GetKeyboardPage() pages.KeyboardPage
	SetReplyID(int)
	SetText(string)
	SetKeyboardPage(pages.KeyboardPage)
}
