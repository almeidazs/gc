package commit

import (
	"github.com/almeidazs/gc/internal/ai"
	"github.com/charmbracelet/huh"
)

func generateMessage(diff string) (string, error) {
	content, err := ai.Prompt(diff)

	if err != nil {
		return "", err
	}

	huh.NewInput().Title("Do you want to edit?").Value(&content).Placeholder(content).Run()

	return content, nil
}
