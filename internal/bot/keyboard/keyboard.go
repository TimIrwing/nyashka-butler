package keyboard

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard/pages"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetKeyboard(p types.KeyboardPage) *tgbotapi.InlineKeyboardMarkup {
	switch p {
	case pages.Config:
		return pages.GetConfig()
	default:
		return nil
	}
}
