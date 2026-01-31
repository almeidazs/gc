package commit

import (
	"github.com/almeidazs/gc/internal/ai"
	"github.com/charmbracelet/huh"
)

func generateMessage(diff string, skip bool, emojis bool) (string, error) {
	content, err := ai.Prompt(diff, emojis)

	if err != nil {
		return "", err
	}

	if !skip {
		huh.NewInput().Title("Do you want to edit?").Value(&content).Placeholder(content).Run()
	}

	return content, nil
}
