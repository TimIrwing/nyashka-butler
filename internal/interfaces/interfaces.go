package interfaces

import (
	"github.com/TimIrwing/nyashka-butler/internal/types"
)

type Message interface {
	New(string) Message
	GetID() int
	GetReplyID() int
	GetChatID() int64
	GetText() string
	GetKeyboardPage() types.KeyboardPage // TODO maybe separate keyboard from message
	SetReplyID(int)
	SetText(string)
	SetKeyboardPage(types.KeyboardPage)
	GetTextInteraction() types.TextInteractionName
}

type Settings interface {
	GetTextInteraction(types.TextInteractionName) *types.SettingsTextInteraction
}