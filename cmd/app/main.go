package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	telegram_bot "goTgExample/pkg/telegram-bot"
	"log"
)

func main() {
	//yandex_gpt.SendMessage()

	bot, err := tgbotapi.NewBotAPI("7112358443:AAGEHc7znXEpAXMyOpk1AKaKlIjE9MbjJ5M")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	tegramBot := telegram_bot.NewBot(bot)
	if err := tegramBot.Start(); err != nil {
		log.Panic(err)
	}

}
