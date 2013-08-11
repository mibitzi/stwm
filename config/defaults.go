package config

func (config *Config) setDefaults() {
	config.Keybinds["mod4-return"] = &Command{
		Cmd:  "exec",
		Args: []string{"urxvt"},
	}
}
