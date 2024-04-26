package yandex_gpt

const (
	foundationModelsUrl = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}
