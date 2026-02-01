package profiles

import (
	"fmt"
	"github.com/almeidazs/gc/internal/config"
	"strings"
)

const (
	reset    = "\033[0m"
	lavender = "\033[38;5;183m"
)

func List() error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	if len(cfg.Profiles) == 0 {
		fmt.Println("Currently you do not have any profile, use \"gc add\"")

		return nil
	}

	var builder strings.Builder

	builder.Grow(len(cfg.Profiles) * 50)

	for name, info := range cfg.Profiles {
		if name == cfg.Current {
			builder.WriteString(lavender)
			builder.WriteString("* ")
			builder.WriteString(reset)
		}

		builder.WriteString(name)
		builder.WriteString(" (")
		builder.WriteString(info.Provider)
		builder.WriteString(") - ")
		builder.WriteString(info.Model)
		builder.WriteByte('\n')
	}

	fmt.Print(builder.String())

	return nil
}
