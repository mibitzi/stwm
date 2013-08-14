package config

func (config *Config) setDefaults() {
	config.Keybinds["mod4-return"] = "exec urxvt"
	config.Keybinds["mod4-shift-h"] = "move left"
	config.Keybinds["mod4-shift-j"] = "move down"
	config.Keybinds["mod4-shift-k"] = "move up"
	config.Keybinds["mod4-shift-l"] = "move right"

	config.Vars["borderWidth"] = "2"
	config.Vars["activeWindowColor"] = "0x0000ff"
	config.Vars["inactiveWindowColor"] = "0x000000"
}
