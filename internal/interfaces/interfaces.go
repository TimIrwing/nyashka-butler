package interfaces

type Message interface {
	New(string) Message
	GetID() int
	GetReplyID() int
	GetChatID() int64
	GetText() string
	SetReplyID(int)
	SetText(string)
}
