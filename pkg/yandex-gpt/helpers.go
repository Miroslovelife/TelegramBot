package yandex_gpt

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
