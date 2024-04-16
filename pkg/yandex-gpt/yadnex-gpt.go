package yandex_gpt

const (
	foundationModelsUrl = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
	model               = "gpt://b1gbn97aajdvpb1a43l2/yandexgpt-lite"
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
