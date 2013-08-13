package config

import (
	"errors"
	"log"
	"strings"
)

// parseContents parses the contents of a config file and changes the
// configuration accordingly.
func (config *Config) parseContents(contents []byte) error {
	lines := strings.Split(string(contents), "\n")

	for _, line := range lines {
		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

		if err := config.parseParts(parts); err != nil {
			log.Printf("Error parsing line: %s\n", line)
			log.Print(err)
		}
	}

	return nil
}

// parseParts parses one line split in to two parts, the type of configuration
// and its value.
func (config *Config) parseParts(parts []string) error {
	if len(parts) != 2 {
		return errors.New("config.parseParts: lines must have two parts")
	}

	parser := make(map[string]func(string) error)
	parser["keybind"] = config.parseKeybind

	if fn, ok := parser[parts[0]]; ok {
		if err := fn(parts[1]); err != nil {
			return err
		}
	} else {
		config.Vars[parts[0]] = parts[1]
	}

	return nil
}

// parseKeybind parses a keybind line. The passed string should have two parts
// separated by a space like:
//     Mod4-Shift-q kill
func (config *Config) parseKeybind(keybind string) error {
	parts := strings.SplitN(keybind, " ", 2)
	if len(parts) != 2 {
		return errors.New("config.parseKeybind: keybinds must have two parts")
	}

	config.Keybinds[strings.ToLower(parts[0])] = parts[1]

	return nil
}
