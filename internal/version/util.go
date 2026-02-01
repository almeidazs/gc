package version

import (
	"context"

	"github.com/almeidazs/gc/internal/exceptions"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func CheckUpdate(ctx context.Context, crr string) (*selfupdate.Release, bool, error) {
	slug := "almeidazs/gc"

	latest, found, err := selfupdate.DetectLatest(slug)

	if err != nil {
		return nil, false, exceptions.InternalError("%v", err)
	}

	if !found || crr == "" {
		return nil, false, nil
	}

	current, err := semver.Parse(crr)

	if err != nil {
		return nil, false, exceptions.InternalError("%v", err)
	}

	if latest.Version.LTE(current) {
		return nil, false, nil
	}

	return latest, true, nil
}
