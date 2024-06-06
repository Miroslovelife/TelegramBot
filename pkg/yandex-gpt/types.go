package yandex_gpt

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

// Response struct for unmarshal response from TgBot
type Response struct {
	Result Result `json:"result"`
}

type Result struct {
	Alternatives []Alternatives `json:"alternatives"`
	Usage        `json:"usage"`
}

type Alternatives struct {
	Message Message `json:"message"`
	Status  string  `json:"status"`
}

type Usage struct {
	InputTextTokens   string `json:"inputTextTokens"`
	CompetitionTokens string `json:"competitionTokens"`
	TotalTokens       string `json:"totalTokens"`
}
