package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseContents(t *testing.T) {
	contents := "keybind Mod4-return exec urxvt\n" +
		"keybind Mod4-q kill\n" +
		"activeWindowColor 0xffffff\n" +
		"font Droid Sans Mono\n"

	config := New()
	assert.NoError(t, config.parseContents([]byte(contents)), "Parse contents")
	assert.Equal(t, len(config.Keybinds), 2, "Number of keybinds parsed")

	firstKey, ok := config.Keybinds["mod4-return"]
	assert.True(t, ok, "First keybind")
	assert.Equal(t, firstKey.Cmd, "exec", "First cmd")
	assert.Equal(t, firstKey.Args[0], "urxvt", "First args")

	secondKey, ok := config.Keybinds["mod4-q"]
	assert.True(t, ok, "Second keybind")
	assert.Equal(t, secondKey.Cmd, "kill", "Second cmd")
	assert.Equal(t, len(secondKey.Args), 0, "Second args")

	color, err := config.IntVar("activeWindowColor")
	assert.NoError(t, err)
	assert.Equal(t, color, 0xffffff)

	font, err := config.StrVar("font")
	assert.NoError(t, err)
	assert.Equal(t, font, "Droid Sans Mono")
}
