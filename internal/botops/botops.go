package botops

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
}

func NewBot() Bot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Println("Telegram token not found")
		os.Exit(1)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	return Bot{
		Bot:     bot,
		Updates: updates,
	}
}
