package main

import (
	"github.com/mibitzi/stwm/commands"
	"github.com/mibitzi/stwm/config"
	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/events"
	"github.com/mibitzi/stwm/keybind"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/log"
	"github.com/mibitzi/stwm/xgb"
)

func main() {
	var err error

	log.Level = log.DEBUG

	cfg := config.New()

	// Entities
	wm, err := wm.New()
	if err != nil {
		log.Fatal(err)
	}
	defer wm.Destroy()

	// Command and event handler
	commands := commands.New(wm)
	events := events.New(wm, commands)

	// xgb
	xgb, err := xgb.New(events, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer xgb.Destroy()

	// Keybinds
	keys := keybinds.New(xgb.X, events, cfg.Keybinds)
	defer keys.Destroy()

	wm.AddWorkspace(workspace.New("1", tiling.New(xgb.ScreenRect())))

	xgb.EnterMain()
}
