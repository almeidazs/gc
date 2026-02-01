package update

import (
	"context"
	"fmt"
	"os"

	"github.com/almeidazs/gc/internal/exceptions"
	"github.com/almeidazs/gc/internal/version"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func Update() error {
	latest, found, err := version.CheckUpdate(context.TODO(), version.Version)

	if err != nil {
		return exceptions.InternalError("couldn’t check for updates: %w", err)
	}

	if !found {
		fmt.Printf("No updates available (Current v%s)\n", version.Version)

		return nil
	}

	fmt.Printf("Update available: %s\nDownloading...\n", latest.Version)

	exe, err := os.Executable()

	if err != nil {
		return exceptions.InternalError("couldn’t get executable path: %w", err)
	}

	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		return exceptions.InternalError("update failed: %w", err)
	}

	fmt.Printf("Update complete ✨ (%v)", latest.Version)

	return nil
}
