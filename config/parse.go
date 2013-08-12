package config

import (
	"errors"
	"log"
	"strings"
)

func (config *Config) parseContents(contents []byte) error {
	lines := strings.Split(string(contents), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)

		if err := config.parseParts(parts); err != nil {
			log.Printf("Error parsing line: %s\n", line)
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
		if err := fn(parts[1:]); err != nil {
			return err
		}
	} else {
		if err := config.parseVar(parts); err != nil {
			return err
		}
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

func (config *Config) parseVar(parts []string) error {
	if len(parts) < 2 {
		return errors.New("Variables must have a value")
	}

	config.Vars[parts[0]] = strings.Join(parts[1:], " ")

	return nil
}
