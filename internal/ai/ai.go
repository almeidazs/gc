package ai

import (
	"github.com/almeidazs/gc/internal/exceptions"
)

func Prompt(diff string, emojis bool) (string, error) {
	key, profile, err := getAI()

	if err != nil {
		return "", err
	}

	prompt := "Analyze the following git diff and generate a commit message. " +
		"The message must follow the Conventional Commits standard. " +
		"Your response should contain *only* the commit message, without any additional text, explanations, or markdown formatting. " +
		"Focus on the primary purpose of the changes and be concise. " +
		"Do not include file names, line numbers, or the diff itself in the output. "

	if emojis || profile.UseEmojis {
		prompt += "Use emojis in the commit message."
	}

	prompt += "Here is the diff:\n" + diff

	return request(key, profile.Provider, profile.Model, prompt)
}

func request(key, provider, model, prompt string) (string, error) {
	switch provider {
	case "openai":
		return requestOpenAI(key, model, prompt)
	case "anthropic":
		return requestAnthropic(key, model, prompt)
	case "google":
		return requestGoogle(key, model, prompt)
	case "meta":
		return requestMeta(key, model, prompt)
	case "xai":
		return requestXAI(key, model, prompt)
	default:
		return "", exceptions.CommandError("unsupported provider: %s", provider)
	}
}
