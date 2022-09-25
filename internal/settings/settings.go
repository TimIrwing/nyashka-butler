package settings

import (
	"errors"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

var ErrIsNil = errors.New("settings: can't save nil value")

type Settings struct {
	col *mongodb.Collection
}

func New(db *mongodb.DB, chatID int64) *Settings {
	return &Settings{col: db.GetChatCollection(chatID)}
}

func (s Settings) GetGeneral() *General {
	p := &General{}
	get := s.col.Get(bson.D{{"isgeneral", true}})
	typed, ok := get.Decode(p).(*General)
	if ok {
		return typed
	}
	return &General{}
}

func (s Settings) GetTextInteraction(name string) *TextInteraction {
	p := &TextInteraction{}
	get := s.col.Get(bson.D{{"TextInteractionName", name}})
	typed, ok := get.Decode(p).(*TextInteraction)
	if ok {
		return typed
	}
	return &TextInteraction{}
}

func (s Settings) SetGeneral(input *General) error {
	if input == nil {
		return ErrIsNil
	}
	input.IsGeneral = true
	return s.col.Add(input)
}

func (s Settings) SetTextInteraction(input *TextInteraction) error {
	if input == nil {
		return ErrIsNil
	}
	return s.col.Add(input)
}
