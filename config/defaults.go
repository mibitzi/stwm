package config

func (config *Config) setDefaults() {
	config.Keybinds["mod4-return"] = &Command{
		Cmd:  "exec",
		Args: []string{"urxvt"},
	}

	config.Vars["borderWidth"] = "2"
	config.Vars["activeWindowColor"] = "0x0000ff"
	config.Vars["inactiveWindowColor"] = "0x000000"
}
