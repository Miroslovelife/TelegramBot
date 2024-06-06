package yandex_gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func newPrompt(model string, options CompletionOptions, messages []Message) *Prompt {
	return &Prompt{
		ModelUri:          model,
		CompletionOptions: options,
		Messages:          messages,
	}
}

func newOption(stream bool, temperature float32, maxTokens string) CompletionOptions {
	return CompletionOptions{
		Stream:      stream,
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}
}

func newMessage(role string, text string) Message {
	return Message{
		Role: role,
		Text: text,
	}
}

func (c *Client) createPrompt(msgs []Message) *Prompt {
	options := newOption(false, 0, "2000")
	messages := msgs
	prompt := newPrompt(c.idCatalog, options, messages)
	return prompt

}

func (c *Client) sendPromptToYandex(prompt []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(prompt))
	if err != nil {
		log.Panic()
	}

	req.Header.Set("Authorization", c.apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal()
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

	err := json.Unmarshal(responseJson, &response)
	if err != nil {
		fmt.Println("error:", err)
	}

	return response, nil
}
