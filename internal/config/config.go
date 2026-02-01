package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const (
	configDir  = ".config/gc"
	configFile = "config.json"
)

var (
	cachedPath string
	pathOnce   sync.Once
)

func GetCurrent() (Profile, error) {
	cfg, err := Load()

	if err != nil {
		return Profile{}, err
	}

	if cfg.Current == "" {
		return Profile{}, fmt.Errorf("no active profile set")
	}

	profile, ok := cfg.Profiles[cfg.Current]

	if !ok {
		return Profile{}, fmt.Errorf("active profile '%s' not found", cfg.Current)
	}

	return profile, nil
}

func Load() (*Config, error) {
	path, err := getPath()

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)

	if os.IsNotExist(err) {
		return &Config{Profiles: make(map[string]Profile, 4)}, nil
	}

	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Profiles == nil {
		cfg.Profiles = make(map[string]Profile, 4)
	}

	return &cfg, nil
}

func getPath() (string, error) {
	pathOnce.Do(func() {
		home, err := os.UserHomeDir()

		if err == nil {
			cachedPath = filepath.Join(home, configDir, configFile)
		}
	})

	if cachedPath == "" {
		home, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}

		return filepath.Join(home, configDir, configFile), nil
	}

	return cachedPath, nil
}
