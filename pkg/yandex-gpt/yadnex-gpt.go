package yandex_gpt

import (
	"log"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	idCatalog  string
}

func NewClient(ApiKey string, IdCatalog string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://llm.api.cloud.yandex.net/foundationModels/v1/completion",
		apiKey:     ApiKey,
		idCatalog:  IdCatalog,
	}
}

func (c *Client) SendMessage(msgs []Message) string {
	prompt := c.createPrompt(msgs)
	promptJson, err := promptToJson(prompt)
	if err != nil {
		log.Fatal("Can't marshal prompt to json")
	}

	response, err := c.sendPromptToYandex(promptJson)
	if err != nil {
		log.Fatal("Can't send request to Yandex API")
	}

	responseJson, err := responseUnmarshal(response)
	if err != nil {
		log.Fatal("Can't unmarshal response to json")
	}

	responseText := responseJson.Result.Alternatives[0].Message.Text
	return responseText
}
