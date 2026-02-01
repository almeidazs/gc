package profiles

import (
	"fmt"

	"github.com/almeidazs/gc/internal/config"
	"github.com/almeidazs/gc/internal/keyring"
	"github.com/almeidazs/gc/internal/style"
	"github.com/charmbracelet/huh"
)

func Add(name, key string) error {
	cfg, err := config.Load()

	if err != nil {
		return err
	}

	if _, exists := cfg.Profiles[name]; exists {
		return fmt.Errorf("\"%s\" is already a profile", name)
	}

	key, provider, model := askModel(key)

	if err := keyring.Set(name, key); err != nil {
		return err
	}

	if err := cfg.Add(config.Profile{
		Name:              name,
		Model:             model,
		UseEmojis:         false,
		AlwaysPush:        false,
		Provider:          provider,
		AlwaysSetUpstream: false,
	}); err != nil {
		return err
	}

	if err := cfg.Save(); err != nil {
		return err
	}

	fmt.Printf("We created and switched to \"%s\"\n", name)

	return nil
}

func askModel(key string) (apikey string, name string, aimodel string) {
	fields := make([]huh.Field, 0, 3)

	if key == "" {
		fields = append(fields,
			huh.NewInput().
				Title("What is the API key?").
				Value(&key).
				EchoMode(huh.EchoModePassword),
		)
	}

	var ai, model string

	fields = append(fields,
		huh.NewSelect[string]().
			Title("Which AI do you want to use?").
			Value(&ai).
			Options(
				huh.NewOption("OpenAI", "openai"),
				huh.NewOption("Anthropic (Claude)", "anthropic"),
				huh.NewOption("Google (Gemini)", "google"),
				huh.NewOption("Meta (LLaMA)", "meta"),
				huh.NewOption("xAI", "xai"),
			),

		huh.NewSelect[string]().
			Title("Now you need to choose the model").
			Value(&model).
			OptionsFunc(func() []huh.Option[string] {
				switch ai {
				case "openai":
					return []huh.Option[string]{
						huh.NewOption("GPT-4o", "gpt-4o"),
						huh.NewOption("o3-mini", "o3-mini"),
						huh.NewOption("GPT-3.5 Turbo", "gpt-3.5-turbo"),
					}
				case "anthropic":
					return []huh.Option[string]{
						huh.NewOption("Claude 3 Opus", "claude-3-opus"),
						huh.NewOption("Claude 3 Sonnet", "claude-3-sonnet"),
						huh.NewOption("Claude 3 Haiku", "claude-3-haiku"),
					}
				case "google":
					return []huh.Option[string]{
						huh.NewOption("Gemini 1.5 Pro", "gemini-1.5-pro"),
						huh.NewOption("Gemini 1.5 Flash", "gemini-1.5-flash"),
						huh.NewOption("Gemini Ultra", "gemini-ultra"),
					}
				case "meta":
					return []huh.Option[string]{
						huh.NewOption("LLaMA 3 70B", "llama-3-70b"),
						huh.NewOption("LLaMA 3 8B", "llama-3-8b"),
						huh.NewOption("Code LLaMA", "code-llama"),
					}
				case "xai":
					return []huh.Option[string]{
						huh.NewOption("Grok-2", "grok-2"),
						huh.NewOption("Grok-1.5", "grok-1.5"),
						huh.NewOption("Grok-Vision", "grok-vision"),
					}
				}

				return nil
			}, &ai),
	)

	huh.NewForm(
		huh.NewGroup(fields...),
	).WithAccessible(style.USE_ACCESSIBLE_MODE).Run()

	return key, ai, model
}
