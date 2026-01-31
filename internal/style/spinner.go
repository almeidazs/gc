package style

import (
	"context"

	"github.com/charmbracelet/huh/spinner"
)

type Spinner struct {
	s      *spinner.Spinner
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSpinner(title string) *Spinner {
	ctx, cancel := context.WithCancel(context.Background())

	s := spinner.New().
		Title(title).
		Context(ctx).
		Accessible(USE_ACCESSIBLE_MODE)

	go s.Run()

	return &Spinner{
		s:      s,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Spinner) Update(title string) {
	s.s.Title(title)
}

func (s *Spinner) Stop() {
	s.cancel()
}
