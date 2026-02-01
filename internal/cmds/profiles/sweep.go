package profiles

import (
	"fmt"

	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/exceptions"
	"github.com/almeidazs/gc/internal/keyring"
)

func Sweep() error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	if cfg.Current == "" {
		return exceptions.CommandError("there are no profiles to sweep")
	}

	if err := keyring.Sweep(); err != nil {
		return exceptions.InternalError("%v", err)
	}

	if err := cfg.Sweep(); err != nil {
		return err
	}

	fmt.Println("Sweeped all profiles")

	return nil
}
