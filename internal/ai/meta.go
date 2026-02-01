package ai

import (
	"encoding/json"

	"github.com/almeidazs/gc/internal/exceptions"
)

const metaURL = "https://api.together.xyz/v1/chat/completions"

type metaRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type metaResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
}

func requestMeta(key, model, prompt string) (string, error) {
	if model == "" {
		model = "meta-llama/Llama-3-70b-chat-hf"
	}

	payload := metaRequest{
		Model:    model,
		Messages: []message{{Role: "user", Content: prompt}},
	}

	body, err := doRequest(metaURL, "POST", "Authorization", "Bearer "+key, payload)

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	var result metaResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	if len(result.Choices) == 0 {
		return "", exceptions.InternalError("no response from meta")
	}

	return result.Choices[0].Message.Content, nil
}
