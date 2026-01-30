package ai

import (
	"encoding/json"
	"fmt"
)

const xaiURL = "https://api.x.ai/v1/chat/completions"

type xaiRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type xaiResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
}

func requestXAI(key, model, prompt string) (string, error) {
	if model == "" {
		model = "grok-2-latest"
	}

	payload := xaiRequest{
		Model:    model,
		Messages: []message{{Role: "user", Content: prompt}},
	}

	body, err := doRequest(xaiURL, "POST", "Authorization", "Bearer "+key, payload)
	if err != nil {
		return "", err
	}

	var result xaiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from xai")
	}

	return result.Choices[0].Message.Content, nil
}
