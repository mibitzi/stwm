package config

func (config *Config) setDefaults() {
	config.Keybinds["mod4-return"] = "exec urxvt"

	config.Vars["borderWidth"] = "2"
	config.Vars["activeWindowColor"] = "0x0000ff"
	config.Vars["inactiveWindowColor"] = "0x000000"
}
