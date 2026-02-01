package commit

import (
	"fmt"

	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/git"
	"github.com/almeidazs/gc/internal/style"
)

func Commit(options CommitOptions) error {
	if err := validateOptions(options); err != nil {
		return err
	}

	spinner := style.NewSpinner("Staging files...")

	defer spinner.Stop()

	if err := git.Stage(options.Files, spinner); err != nil {
		return err
	}

	spinner.Update("Currently analyzing the changes to generate the message...")

	diff, err := git.StagedDiff()

	if err != nil {
		return err
	}

	spinner.Update("Fetching your current profile...")

	profile, err := config.GetCurrent()

	if err != nil {
		return err
	}

	spinner.Update("Trying to generate the message or use the -m one...")

	message, err := resolveMessage(options, profile, diff)

	if err != nil {
		return err
	}

	spinner.Update(fmt.Sprintf("Committing the message (%d chars)...", len(message)))

	if err := git.Commit(message); err != nil {
		return err
	}

	if options.Push || profile.AlwaysPush {
		setUpstream := options.SetUpstream || profile.AlwaysSetUpstream
		message := "Automatic push detected, trying to push to the branch"

		if setUpstream {
			message += " and setting the upstream"
		}

		spinner.Update(message)

		return push(PushOptions{
			Branch:      options.Branch,
			SetUpstream: setUpstream,
		})
	}

	return nil
}
