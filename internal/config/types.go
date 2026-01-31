package config

type Profile struct {
	// TODO: Add a command to update the profile options

	UseEmojis  bool   `json:"emojis"`
	AlwaysPush bool   `json:"always_push"`
	Name       string `json:"name"`
	Provider   string `json:"provider"`
	Model      string `json:"model,omitempty"`
}

type Config struct {
	Current  string             `json:"current"`
	Profiles map[string]Profile `json:"profiles"`
}
