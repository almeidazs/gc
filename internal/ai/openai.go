package ai

import (
	"encoding/json"
	"fmt"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

type openaiRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type openaiResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
}

func requestOpenAI(key, model, prompt string) (string, error) {
	if model == "" {
		model = "gpt-4o"
	}

	payload := openaiRequest{
		Model:    model,
		Messages: []message{{Role: "user", Content: prompt}},
	}

	body, err := doRequest(openaiURL, "POST", "Authorization", "Bearer "+key, payload)
	if err != nil {
		return "", err
	}

	var result openaiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from openai")
	}

	return result.Choices[0].Message.Content, nil
}
