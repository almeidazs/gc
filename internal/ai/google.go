package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/almeidazs/gc/internal/exceptions"
)

const googleURL = "https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s"

type googleRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type googleResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func requestGoogle(key, model, prompt string) (string, error) {
	if model == "" {
		model = "gemini-1.5-flash"
	}

	url := fmt.Sprintf(googleURL, model, key)

	payload := googleRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{Parts: []struct {
				Text string `json:"text"`
			}{{Text: prompt}}},
		},
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))

	if err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	req.Header.Set("Content-Type", "application/json")

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

	var result googleResponse

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", exceptions.InternalError("no response from google")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
