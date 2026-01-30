package config

type Profile struct {
	Name	string `json:"name"`
	Provider string `json:"provider"`
	Model	string `json:"model,omitempty"`
}

type Config struct {
	Current string	`json:"current"`
	Profiles map[string]Profile `json:"profiles"`
}

func (c *Config) Add(profile Profile) {
	c.Current = profile.Name
	c.Profiles[profile.Name] = profile
}

func (c *Config) Remove(name string) {
	delete(c.Profiles, name)
}
