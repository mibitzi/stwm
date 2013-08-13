package config

import (
	"fmt"
	"strconv"
)

func (config *Config) GetVar(name string) (string, error) {
	if val, ok := config.Vars[name]; ok {
		return val, nil
	} else {
		return "", fmt.Errorf("config.GetVar: \"%s\" not found", name)
	}
}

// IntVar returns a config variable as an int.
func (config *Config) IntVar(name string) (int, error) {
	if val, err := config.GetVar(name); err != nil {
		return 0, err
	} else {
		if i, err := strconv.ParseInt(val, 0, 32); err != nil {
			return 0, fmt.Errorf("config.IntVar: \"%s\": %s", name, err)
		} else {
			return int(i), nil
		}
	}
}

// StrVar returns a config variable as a string.
func (config *Config) StrVar(name string) (string, error) {
	if val, err := config.GetVar(name); err != nil {
		return "", err
	} else {
		return val, nil
	}
}
