package commit

import (
	"os"

	"github.com/almeidazs/gc/internal/ai"
	"github.com/almeidazs/gc/internal/style"
	"github.com/charmbracelet/huh"
)

func generateMessage(diff string, skip bool, emojis bool) (string, error) {
	content, err := ai.Prompt(diff, emojis)
	if err != nil {
		return "", err
	}

	if skip {
		return content, nil
	}

	input := huh.NewInput().
		Title("Do you want to edit?").
		Value(&content).
		Placeholder(content)

	if style.USE_ACCESSIBLE_MODE {
		if err := input.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return "", err
		}

		return content, nil
	}

	if err := input.Run(); err != nil {
		return "", err
	}

	return content, nil
}
