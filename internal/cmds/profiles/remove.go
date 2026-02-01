package profiles

import (
	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/exceptions"
	"github.com/almeidazs/gc/internal/keyring"
)

func Remove(name string) error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	isCrr := cfg.Current == name

	if isCrr {
		return exceptions.CommandError("currently \"%s\" is your profile, you need to switch it first", name)
	}

	if _, exists := cfg.Profiles[name]; !exists {
		return exceptions.CommandError("we could not found the profile \"%s\"\n", name)
	}

	if err := keyring.Remove(name); err != nil {
		return exceptions.InternalError("%v", err)
	}

	if err := cfg.Remove(name); err != nil {
		return err
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	return nil
}
