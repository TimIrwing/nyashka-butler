package settings

import (
	"errors"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	"go.mongodb.org/mongo-driver/bson"
)

var ErrIsNil = errors.New("settings: can't save nil value")

type Settings struct {
	col *mongodb.Collection
}

func New(db *mongodb.DB, chatID int64) *Settings {
	return &Settings{col: db.GetChatCollection(chatID)}
}

func (s Settings) GetGeneral() *types.SettingsGeneral {
	p := &types.SettingsGeneral{}
	get := s.col.Get(bson.D{{"isgeneral", true}})
	typed, ok := get.Decode(p).(*types.SettingsGeneral)
	if ok {
		return typed
	}
	return &types.SettingsGeneral{}
}

func (s Settings) GetTextInteraction(name types.TextInteractionName) *types.SettingsTextInteraction {
	p := &types.SettingsTextInteraction{}
	get := s.col.Get(bson.D{{"textinteractionname", name}})
	typed, ok := get.Decode(p).(*types.SettingsTextInteraction)
	if ok {
		return typed
	}
	return &types.SettingsTextInteraction{}
}

func (s Settings) SetGeneral(input *types.SettingsGeneral) error {
	if input == nil {
		return ErrIsNil
	}
	input.IsGeneral = true
	return s.col.Add(input)
}

func (s Settings) SetTextInteraction(input *types.SettingsTextInteraction) error {
	if input == nil {
		return ErrIsNil
	}
	return s.col.Add(input)
}
