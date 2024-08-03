package botops

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) GetMainMenu() tgbotapi.InlineKeyboardMarkup {

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
			tgbotapi.NewInlineKeyboardButtonData("2 🧲", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("jokes", "jokes"),
		),
	)
	return numericKeyboard
}
