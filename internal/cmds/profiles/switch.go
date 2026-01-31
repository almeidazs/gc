package profiles

import (
	"github.com/almeidazs/gc/internal/config"
)

func Switch(name string) error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	return cfg.Switch(name)
}
