package ai

import (
	"fmt"
	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/keyring"
)

func getAI() (string, config.Profile, error) {
	current, err := config.GetCurrent()
	if err != nil {
		return "", config.Profile{}, err
	}

	key, err := keyring.Get(current.Name)
	if err != nil {
		return "", config.Profile{}, fmt.Errorf("failed to get API key: %w", err)
	}

	return key, current, nil
}
