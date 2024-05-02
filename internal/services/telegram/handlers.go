package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	yandex_gpt "goTgExample/pkg/yandex-gpt"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды нет")
	switch message.Command() {
	case commandStart:
		msg.Text = "Привет это gpt бот \n Для начачла выбирете модель AI\n/yandex \n/gpt(Пока недоступно)"
		_, err := b.bot.Send(msg)
		return err
	case commandChoiceYandexGPT:
		aiState = commandChoiceYandexGPT
		msg.Text = "Отлично. Вы выбрали модель YandexGPT, можете делать запросы"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err

	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message, aiState string) {
	log.Printf("[%s]", message.Text)
	if aiState == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Необходимо выбрать AI модель")
		b.bot.Send(msg)
	}
	switch aiState {
	case commandChoiceYandexGPT:
		yandexGptResponse := yandex_gpt.GetResponseText(message.Text)

		msg := tgbotapi.NewMessage(message.Chat.ID, yandexGptResponse)
		msg.ParseMode = "Markdown"
		b.bot.Send(msg)
	}

}
