package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"goTgExample/internal/models"
	yandex_gpt "goTgExample/pkg/yandex-gpt"
	"log"
	"os"
)

func (b *Bot) handleCommand(message *tgbotapi.Message, userID int) error {

	msg := tgbotapi.NewMessage(message.Chat.ID, "Такой команды нет")
	switch message.Command() {
	case commandStart:
		b.VerifyUser(context.Background(), userID)
		msg.Text = "Привет это gpt бот \n Для начачла выбирете модель AI\n/yandex \n/gpt(Пока недоступно)"
		_, err := b.bot.Send(msg)

		return err
	case commandChoiceYandexGPT:
		aiState = commandChoiceYandexGPT
		msg.Text = "Отлично. Вы выбрали модель YandexGPT, можете делать запросы."
		_, err := b.bot.Send(msg)
		return err
	case commandResetContext:
		b.ResetContext(context.Background(), userID)
		msg.Text = "Вы сбросили контекст беседы."
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err

	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message, aiState string, userID int) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	yandexApiKey, yandexIdCatalog := os.Getenv("YANDEX_API_KEY"), os.Getenv("YANDEX_ID_CATALOG")
	yandexClient := yandex_gpt.NewClient(yandexApiKey, yandexIdCatalog)

	log.Printf("[%s]", message.Text)
	if aiState == "" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Необходимо выбрать AI модель")
		b.bot.Send(msg)
	}
	switch aiState {
	case commandChoiceYandexGPT:
		userMessage := yandex_gpt.Message{
			Role: "user",
			Text: message.Text,
		}

		chatHistory, err := b.GetChatHistory(context.Background(), userID)
		if err != nil {
			fmt.Errorf("error get history chat: %s", err)
		}

		chat := append(*chatHistory, userMessage)
		fmt.Println(chat)
		yandexGptResponse := yandexClient.SendMessage(chat)

		Msg := models.Message{UserMessage: userMessage.Text, BotMessage: yandexGptResponse}
		err = b.NewMessage(context.Background(), userID, Msg)
		if err != nil {
			fmt.Errorf("can't save message: %s", err)
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, yandexGptResponse)
		msg.ParseMode = "Markdown"
		b.bot.Send(msg)
	}

}

func (b *Bot) ResetContext(ctx context.Context, userID int) error {
	return b.us.ResetContext(ctx, userID)
}

func (b *Bot) VerifyUser(ctx context.Context, userID int) error {
	return b.us.AddUser(ctx, userID)
}

func (b *Bot) GetChatHistory(ctx context.Context, userID int) (msgs *[]yandex_gpt.Message, err error) {
	chat, err := b.us.GetChatHistory(ctx, userID)
	if err != nil {
		return nil, err
	}

	messages := chat.Messages

	yandexMsgs := []yandex_gpt.Message{}
	for _, msg := range messages {
		// Добавление UserMessage в yandexMessages
		if msg.UserMessage != "" {
			yandexMsgs = append(yandexMsgs, yandex_gpt.Message{
				Role: "user",
				Text: msg.UserMessage,
			})
		}

		// Добавление BotMessage в yandexMessages
		if msg.BotMessage != "" {
			yandexMsgs = append(yandexMsgs, yandex_gpt.Message{
				Role: "assistant",
				Text: msg.BotMessage,
			})
		}
	}

	return &yandexMsgs, nil
}

func (b *Bot) NewMessage(ctx context.Context, userId int, message models.Message) error {
	return b.us.NewMessage(ctx, userId, message)
}
