package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/almeidazs/gc/internal/exceptions"
)

var httpClient = &http.Client{
	Timeout: 60 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func doRequest(url, method, authHeader, authValue string, payload interface{}) ([]byte, error) {
	body, err := json.Marshal(payload)

	if err != nil {
		return nil, exceptions.InternalError("%v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		return nil, exceptions.InternalError("%v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(authHeader, authValue)

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, exceptions.InternalError("%v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, exceptions.InternalError("%v", err)
	}

	if resp.StatusCode != 200 {
		return nil, exceptions.InternalError("http %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}
