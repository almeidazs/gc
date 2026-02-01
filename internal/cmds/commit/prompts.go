package commit

import (
	"strings"

	"github.com/almeidazs/gc/internal/exceptions"
	"github.com/almeidazs/gc/internal/style"
	"github.com/charmbracelet/huh"
)

func askCoauthor() (string, string, error) {
	var name, email string

	var form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the name of the co-author?").
				Value(&name),
			huh.NewInput().
				Title("And now what is the email?").
				Value(&email).
				SuggestionsFunc(func() []string {
					if email == "" || strings.Contains(email, "@") {
						return []string{}
					}

					return []string{email + "@gmail.com", email + "@outlook.com"}
				}, &email),
		),
	).WithAccessible(style.USE_ACCESSIBLE_MODE)

	if err := form.Run(); err != nil {
		return "", "", exceptions.InternalError("%v", err)
	}

	return name, email, nil
}
