package main

import (
	"log"
	"tg-devops-bot/internal/botops"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot := botops.NewBot()

	log.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	for update := range bot.Updates {
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {

			case "/start":
				msg.ReplyMarkup = bot.GetMainMenu()
			}

			if _, err := bot.Bot.Send(msg); err != nil {
				panic(err)
			}

		} else if update.CallbackQuery != nil {

			bot.GeneralHandler(update)

		}
	}
}
