package yandex_gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func handleTelegramText(text string) Response {
	msg := newMessage("user", text)
	prompt := createPrompt(msg)

	promptJson, err := promptToJson(prompt)
	if err != nil {
		log.Fatal("Can't marshal prompt to error:", err)
	}

	responseJsonMessage, err := sendPromptToYandex(promptJson)
	if err != nil {
		log.Fatal("sendPrompt error:", err)
	}

	responseUnmarshalMessage, err := responseUnmarshal(responseJsonMessage)

	return responseUnmarshalMessage
}

func createPrompt(msg Message) *Prompt {
	options := newOption(false, 0, "2000")
	messages := []Message{}
	messages = append(messages, msg)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	yandexIdCatalog := os.Getenv("YANDEX_ID_CATALOG")
	prompt := newPrompt(yandexIdCatalog, options, messages)
	return prompt
}

func sendPromptToYandex(prompt []byte) ([]byte, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	yandexApi := os.Getenv("YANDEX_API_KEY")
	req, err := http.NewRequest("POST", foundationModelsUrl, bytes.NewBuffer(prompt))
	if err != nil {
		log.Panic()
	}

	req.Header.Set("Authorization", yandexApi)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic()
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return body, nil

}

func promptToJson(prompt *Prompt) ([]byte, error) {
	jsonPrompt, err := json.Marshal(prompt)
	if err != nil {
		return nil, err
	}

	return jsonPrompt, nil
}

func responseUnmarshal(responseJson []byte) (Response, error) {
	var response Response

	err := json.Unmarshal([]byte(responseJson), &response)
	if err != nil {
		fmt.Println("error:", err)
	}

	return response, nil
}

func GetResponseText(text string) string {
	response := handleTelegramText(text)
	responseText := response.Result.Alternatives[0].Message.Text

	return responseText
}
