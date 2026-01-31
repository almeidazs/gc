package commit

import (
	"fmt"

	"github.com/almeidazs/gc/internal/git"
)

func Commit(options CommitOptions) error {
	if err := validateOptions(options); err != nil {
		return err
	}

	fmt.Println("Staging files...")

	if err := git.Stage(options.Files); err != nil {
		return err
	}

	fmt.Println("Currently analyzing the changes and generating the message...")

	diff, err := git.StagedDiff()

	if err != nil {
		return err
	}

	message, err := resolveMessage(options, diff)

	if err != nil {
		return err
	}

	fmt.Printf("Committing (%d chars)...\n", len(message))

	if err := git.Commit(message); err != nil {
		return err
	}

	if options.Push {
		return push(options.Branch)
	}

	return nil
}
