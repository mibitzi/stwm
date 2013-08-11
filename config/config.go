package config

import (
	"io/ioutil"
)

type Command struct {
	Cmd  string
	Args []string
}

type Config struct {
	Keybinds map[string]*Command
}

func New() *Config {
	config := &Config{
		Keybinds: make(map[string]*Command),
	}

	config.setDefaults()

	return config
}

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
