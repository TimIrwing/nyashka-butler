package pages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetConfig() *tgbotapi.InlineKeyboardMarkup {
	var kb [][]tgbotapi.InlineKeyboardButton

	kb = append(kb, []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("💜", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Test", "2"),
	})

	res := tgbotapi.InlineKeyboardMarkup{InlineKeyboard: kb}
	return &res
}
