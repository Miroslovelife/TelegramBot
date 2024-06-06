package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"goTgExample/internal/storage"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	us  *storage.UserStorage
}

func NewBot(bot *tgbotapi.BotAPI, userStorage *storage.UserStorage) *Bot {
	return &Bot{bot: bot, us: userStorage}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)
	return nil
}

const (
	commandStart           = "start"
	commandChoiceYandexGPT = "yandex"
	commandResetContext    = "reset"
)

var aiState string

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		userID := update.Message.From.ID
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message, userID)
			continue
		}
		b.handleMessage(update.Message, aiState, userID)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
