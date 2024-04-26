package app

import (
	"goTgExample/internal/services/free_ai"
	"goTgExample/internal/services/telegram"
)

func Run() {
	telegram.StartTelegramBot()
	free_ai.FreeAiStart()
}
