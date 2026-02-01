package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/almeidazs/gc/internal/exceptions"
)

func (c *Config) Save() error {
	path, err := getPath()

	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return exceptions.InternalError("%v", err)
	}

	data, err := json.MarshalIndent(c, "", "\t")

	if err != nil {
		return exceptions.InternalError("%v", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return exceptions.InternalError("%v", err)
	}

	return nil
}

func (c *Config) Add(profile Profile) error {
	if profile.Name == "" {
		return exceptions.CommandError("profile name cannot be empty")
	}

	if profile.Provider == "" {
		return exceptions.CommandError("provider cannot be empty")
	}

	c.Current = profile.Name
	c.Profiles[profile.Name] = profile

	return nil
}

func (c *Config) Remove(name string) error {
	if _, ok := c.Profiles[name]; !ok {
		return exceptions.CommandError("profile '%s' not found", name)
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

func (c *Config) Sweep() error {
	c.Current = ""
	c.Profiles = map[string]Profile{}

	return c.Save()
}

func (c *Config) Switch(name string) error {
	if c.Current == name {
		return exceptions.CommandError("This is already your current profile")
	}

	if _, exists := c.Profiles[name]; !exists {
		return exceptions.CommandError("\"%s\" is not a profile", name)
	}

	if _, ok := c.Profiles[name]; !ok {
		return exceptions.CommandError("profile '%s' not found", name)
	}

	c.Current = name

	return c.Save()
}
