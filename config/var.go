package config

import (
	"fmt"
	"strconv"
)

// IntVar returns a config variable as an int.
func (config *Config) IntVar(name string) (int, error) {
	if val, ok := config.Vars[name]; ok {
		if i, err := strconv.ParseInt(val, 0, 32); err != nil {
			return 0, fmt.Errorf("Variable \"%s\": %s", name, err)
		} else {
			return int(i), nil
		}
	} else {
		return 0, fmt.Errorf("Variable \"%s\" not found", name)
	}
}
