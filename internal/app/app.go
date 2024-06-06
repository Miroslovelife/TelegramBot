package app

import (
	"fmt"
	"goTgExample/internal/config"
	"goTgExample/internal/services/telegram"
	"goTgExample/internal/storage"
)

func Run() {
	cfg, err := config.MustLoadConfig()
	if err != nil {
		fmt.Errorf("error loading config: %v", err)
	}

	pool, err := storage.ConnectDB(cfg.Database)
	if err != nil {
		fmt.Errorf("error connecting to database: %v", err)
	}

	userStorage := storage.NewUserStorage(pool)

	telegram.StartTelegramBot(userStorage)
}
