package config

import "fmt"

type Profile struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Model    string `json:"model,omitempty"`
}

type Config struct {
	Current  string             `json:"current"`
	Profiles map[string]Profile `json:"profiles"`
}

func (c *Config) Add(profile Profile) error {
	if profile.Name == "" {
		return fmt.Errorf("profile name cannot be empty")
	}

	if profile.Provider == "" {
		return fmt.Errorf("provider cannot be empty")
	}

	c.Current = profile.Name
	c.Profiles[profile.Name] = profile

	return nil
}

func (c *Config) Remove(name string) error {
	if _, ok := c.Profiles[name]; !ok {
		return fmt.Errorf("profile '%s' not found", name)
	}

	delete(c.Profiles, name)

	if c.Current == name {
		c.Current = ""

		for pname := range c.Profiles {
			c.Current = pname

			break
		}
	}

	return nil
}

func (c *Config) Switch(name string) error {
	if _, ok := c.Profiles[name]; !ok {
		return fmt.Errorf("profile '%s' not found", name)
	}

	c.Current = name

	return nil
}
