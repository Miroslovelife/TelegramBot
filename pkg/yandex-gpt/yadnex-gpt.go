package yandex_gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	foundationModelsUrl = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
	model               = "gpt://b1gbn97aajdvpb1a43l2/yandexgpt-lite"
)

type Prompt struct {
	ModelUri          string `json:"modelUri"`
	CompletionOptions `json:"completionOptions"`
	Messages          []Message `json:"messages"`
}

type Headers struct {
	Url           string `json:"url"`
	Authorization string `json:"authorization"`
}

type CompletionOptions struct {
	Stream      bool    `json:"stream"`
	Temperature float32 `json:"temperature"`
	MaxTokens   string  `json:"maxTokens"`
}

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

func NewPrompt(model string, options CompletionOptions, messages []Message) *Prompt {
	return &Prompt{
		ModelUri:          model,
		CompletionOptions: options,
		Messages:          messages,
	}
}

func NewHeaders(url string, token string) *Headers {
	return &Headers{
		Url:           url,
		Authorization: token,
	}
}

func NewOption(stream bool, temperature float32, maxTokens string) CompletionOptions {
	return CompletionOptions{
		Stream:      stream,
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}
}

func NewMessage(role string, text string) Message {
	return Message{
		Role: role,
		Text: text,
	}
}

func SendMessage() {
	a := NewOption(false, 0, "2000")
	message := NewMessage("user", "как сделать компьютер")
	messages := []Message{}
	messages = append(messages, message)
	prompt := NewPrompt(model, a, messages)
	header := NewHeaders(foundationModelsUrl, "Api-Key AQVNySxYTJWaH3Kw-h1hHKI3EegY6XVbca-06Bz_")
	promptJson, err := json.Marshal(prompt)
	if err != nil {

	}

	headerJson, err := json.Marshal(header)
	if err != nil {

	}
	fmt.Println("Prompt: \n", string(promptJson), "\n", "Headers: \n", string(headerJson))

	fmt.Println(bytes.NewBuffer(promptJson))

	req, err := http.NewRequest("POST", foundationModelsUrl, bytes.NewBuffer(promptJson))
	if err != nil {
		log.Panic()
	}

	req.Header.Set("Authorization", "Api-Key AQVNySxYTJWaH3Kw-h1hHKI3EegY6XVbca-06Bz_")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение тела ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return
	}

	// Вывод тела ответа
	fmt.Println("Ответ от сервера:", string(body))
}
