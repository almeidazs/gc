package git

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func execCmd(name string, args ...string) (string, error) {
	var output bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(output.String()), nil
}

func CurrentBranch() (string, error) {
	return execCmd("git", "branch", "--show-current")
}

func ResolveBranch(branch string) (string, error) {
	if branch != "" {
		return branch, nil
	}

	return CurrentBranch()
}

func Stage(files []string) error {
	args := []string{"add"}

	if len(files) == 0 {
		args = append(args, ".")

		fmt.Printf("No files provided, staging all...")
	} else {
		args = append(args, files...)
	}

	cmd := exec.Command("git", args...)

	cmd.Stdout = io.Discard

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return fmt.Errorf("git add failed: %s", strings.TrimSpace(stderr.String()))
		}

		return fmt.Errorf("git add failed: %w", err)
	}

	return nil
}

func StagedDiff() (string, error) {
	var output bytes.Buffer

	cmd := exec.Command("git", "diff", "--staged")
	cmd.Stdout = &output
	cmd.Stderr = &output

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(
			"git diff --staged failed: %w\n%s",
			err,
			strings.TrimSpace(output.String()),
		)
	}

	diff := output.String()

	if strings.TrimSpace(diff) == "" {
		return "", errors.New("no staged changes to commit")
	}

	return diff, nil
}

func Commit(message string) error {
	args := []string{"commit", "-m", message}

	cmd := exec.Command("git", args...)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return fmt.Errorf("git commit failed: %s", strings.TrimSpace(stderr.String()))
		}

		return fmt.Errorf("git commit failed: %w", err)
	}

	return nil
}

func Push(branch string) error {
	args := []string{"push", "origin", branch}

	cmd := exec.Command("git", args...)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return fmt.Errorf("git push failed: %s", strings.TrimSpace(stderr.String()))
		}

		return fmt.Errorf("git push failed: %w", err)
	}

	return nil
}
