package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func StartTelegramBot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tgApi := os.Getenv("TELEGRAM_API_KEY")

	bot, err := tgbotapi.NewBotAPI(tgApi)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Panic(err)
	}
}
