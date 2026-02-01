package commit

import (
	"fmt"
	"os"

	"github.com/almeidazs/gc/internal/ai"
	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/exceptions"
	"github.com/almeidazs/gc/internal/git"
	"github.com/almeidazs/gc/internal/style"
	"github.com/charmbracelet/huh"
)

func push(opts PushOptions) error {
	resolved, err := git.ResolveBranch(opts.Branch)

	if err != nil {
		return err
	}

	if err := git.Push(resolved, opts.Force, opts.SetUpstream); err != nil {
		return err
	}

	fmt.Printf("Pushed automatically to the branch \"%s\"\n", resolved)

	return nil
}

func resolveMessage(opts CommitOptions, profile config.Profile, diff string) (string, error) {
	if opts.Message != "" {
		fmt.Printf("Using custom message (%d chars)...\n", len(opts.Message))

		return opts.Message, nil
	}

	msg, err := generateMessage(diff, opts.SkipPrompts, opts.Emojis || profile.UseEmojis)

	if err != nil {
		return "", err
	}

	if !opts.Coauthored {
		return msg, nil
	}

	name, email, err := askCoauthor()

	if err != nil {
		return "", err
	}

	return msg + fmt.Sprintf("\n\nCo-authored-by: %s <%s>", name, email), nil
}

func validateOptions(opts CommitOptions) error {
	if opts.Coauthored && opts.SkipPrompts {
		return exceptions.CommandError("--coauthored cannot be used with --yes")
	}

	return nil
}

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
			return "", exceptions.InternalError("%v", err)
		}

		return content, nil
	}

	if err := input.Run(); err != nil {
		return "", exceptions.InternalError("%v", err)
	}

	return content, nil
}
