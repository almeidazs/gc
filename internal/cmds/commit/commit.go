package commit

import (
	"fmt"

	"github.com/almeidazs/gc/internal/git"
)

func Commit(options CommitOptions) error {
	// Ignore branch for now (Will be used later)
	_, err := git.ResolveBranch(options.Branch)

	if err != nil {
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

	message, err := generateMessage(diff)

	if err != nil {
		return err
	}

	if options.Coauthored {
		name, email, err := askCoauthor()

		if err != nil {
			return err
		}

		message += fmt.Sprintf("\n\nCo-authored-by: %s <%s>", name, email)
	}

	fmt.Printf("Message generated (%v chars), commiting...", len(message))

	if git.Commit(message); err != nil {
		return err
	}

	return nil
}
