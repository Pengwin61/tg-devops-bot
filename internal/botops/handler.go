package botops

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) GeneralHandler(update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

	if _, err := bot.Bot.Request(callback); err != nil {
		panic(err)
	}

	if callback.Text == "jokes" {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "шутка")
		if _, err := bot.Bot.Send(msg); err != nil {
			panic(err)
		}
	} else {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		if _, err := bot.Bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
