package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStrat = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case commandStrat:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввел команду /start")

		b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды нет")

		b.bot.Send(msg)
	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)

}
