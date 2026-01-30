package commit

import (
	"fmt"

	"github.com/almeidazs/gc/internal/git"
)

func Commit(branch string, files []string) error {
	branch, err := git.ResolveBranch(branch)

	if err != nil {
		return err
	}

	fmt.Println("Staging files...")

	if err := git.Stage(files); err != nil {
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

	fmt.Printf("Message generated (%v chars), commiting...", len(message))

	if git.Commit(message); err != nil {
		return err
	}

	return nil
}
