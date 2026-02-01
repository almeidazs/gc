package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/almeidazs/gc/internal/exceptions"
)

const anthropicURL = "https://api.anthropic.com/v1/messages"

type anthropicRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []message `json:"messages"`
}

type anthropicResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func requestAnthropic(key, model, prompt string) (string, error) {
	if model == "" {
		model = "claude-3-sonnet-20240229"
	}

	payload := anthropicRequest{
		Model:     model,
		MaxTokens: 256,
		Messages:  []message{{Role: "user", Content: prompt}},
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	req, err := http.NewRequest("POST", anthropicURL, bytes.NewReader(body))

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", key)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := httpClient.Do(req)

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	if resp.StatusCode != 200 {
		return "", exceptions.InternalError("http %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result anthropicResponse

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", err
	}

	if len(result.Content) == 0 {
		return "", exceptions.InternalError("no response from anthropic")
	}

	return result.Content[0].Text, nil
}
