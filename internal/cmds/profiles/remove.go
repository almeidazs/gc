package profiles

import (
	"fmt"

	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/keyring"
)

func Remove(name string) error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	isCrr := cfg.Current == name

	if isCrr {
		return fmt.Errorf("currently \"%s\" is your profile, you need to switch it first", name)
	}

	if _, exists := cfg.Profiles[name]; !exists {
		return fmt.Errorf("we could not found the profile \"%s\"\n", name)
	}

	if err := keyring.Remove(name); err != nil {
		return err
	}

	cfg.Remove(name)

	if err := config.Save(cfg); err != nil {
		return err
	}

	return nil
}
