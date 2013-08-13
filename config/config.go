package config

import (
	"io/ioutil"
)

type Config struct {
	// Keybinds holds key-sequences and actions which should be taken.
	Keybinds map[string]string

	// Vars holds all configured variables.
	Vars map[string]string
}

// New creates a new Config instance with all default values set.
func New() *Config {
	config := &Config{
		Keybinds: make(map[string]string),
		Vars:     make(map[string]string),
	}

	config.setDefaults()

	return config
}

// Load loads a file and parses its content.
func (config *Config) Load(file string) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = config.parseContents(contents)
	if err != nil {
		return err
	}

	return nil
}
