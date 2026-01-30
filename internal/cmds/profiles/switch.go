package profiles

import (
	"fmt"

	"github.com/almeidazs/gc/internal/config"
)

func Switch(name string) error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	if cfg.Current == name {
		return fmt.Errorf("This is already your current profile")
	}

	if _, exists := cfg.Profiles[name]; !exists {
		return fmt.Errorf("\"%s\" is not a profile", name)
	}

	cfg.Current = name

	if err := config.Save(cfg); err != nil {
		return err
	}

	return nil
}
