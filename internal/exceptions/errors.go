package exceptions

import (
	"errors"
	"fmt"
)

var (
	ErrCommand = errors.New("command err")
	ErrInternal = errors.New("internal")
)

func ExitCode(err error) int {
	switch {
	case err == nil:
		return 0
	case errors.Is(err, ErrCommand):
		return 1
	default:
		return 2
	}
}

func CommandError(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrCommand, fmt.Sprintf(format, args...))
}

func InternalError(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrInternal, fmt.Sprintf(format, args...))
}
