package config

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func (config *Config) parseContents(contents []byte) error {
	lines := strings.Split(string(contents), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)

		if err := config.parseParts(parts); err != nil {
			log.Println("Error parsing line: %s", line)
			log.Print(err)
		}
	}

	return nil
}

func (config *Config) parseParts(parts []string) error {
	if len(parts) == 0 {
		return nil
	}

	parser := make(map[string]func([]string) error)
	parser["keybind"] = config.parseKeybind

	if fn, ok := parser[parts[0]]; ok {
		fn(parts[1:])
	} else {
		return fmt.Errorf("Unknown config type: %s", parts[0])
	}

	return nil
}

func (config *Config) parseKeybind(parts []string) error {
	if len(parts) == 0 {
		return errors.New("No keys specified")
	} else if len(parts) == 1 {
		return errors.New("No command specified")
	}

	command := &Command{
		Cmd:  parts[1],
		Args: parts[2:],
	}

	config.Keybinds[strings.ToLower(parts[0])] = command

	return nil
}
