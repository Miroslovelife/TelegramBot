package models

type User struct {
	UserId string `json:"user_id"`
	Chat   Chat   `json:"chat"`
}

type Chat struct {
	Messages []Message
}

type Message struct {
	BotMessage  string `json:"bot_message"`
	UserMessage string `json:"user_message"`
}
