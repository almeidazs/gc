package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	CONFIG_DIR   = ".config/gc"
    CONFIG_FILE  = "config.json"
)

func Load() (*Config, error) {
    path, err := getPath()

    if err != nil {
        return nil, err
    }

    data, err := os.ReadFile(path)

    if os.IsNotExist(err) {
        return &Config{Profiles: map[string]Profile{}}, nil
    }

    if err != nil {
        return nil, err
    }

    var cfg Config
    
	if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}

func Save(cfg *Config) error {
    path, err := getPath()

    if err != nil {
        return err
    }

    if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
        return err
    }

    data, err := json.MarshalIndent(cfg, "", "\t")

    if err != nil {
        return err
    }

    return os.WriteFile(path, data, 0600)
}

func getPath() (path string, err error) {
	home, err := os.UserHomeDir()

    if err != nil {
        return "", err
    }

    return filepath.Join(home, CONFIG_DIR, CONFIG_FILE), nil
}
