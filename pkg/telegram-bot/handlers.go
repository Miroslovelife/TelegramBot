package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	yandex_gpt "goTgExample/pkg/yandex-gpt"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды нет")
	switch message.Command() {
	case commandStart:
		msg.Text = "Вы ввели команду /start"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err

	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s]", message.Text)
	yandexGptResponse := yandex_gpt.GetResponseText(message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, yandexGptResponse)
	msg.ParseMode = "Markdown"
	b.bot.Send(msg)
}
