package main

import (
	"github.com/mibitzi/stwm/commands"
	"github.com/mibitzi/stwm/config"
	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/events"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/log"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/xgb"
)

func main() {
	var err error

	log.Level = log.DEBUG

	cfg := config.New()

	wm, err := wm.New()
	if err != nil {
		log.Fatal(err)
	}
	defer wm.Destroy()

	cmd := commands.New(wm)
	ev := events.New(wm, cmd)

	xgb, err := xgb.New(ev, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer xgb.Destroy()

	wm.AddWorkspace(workspace.New("1", tiling.New(rect.New(0, 0, 1024, 768))))

	xgb.EnterMain()
}
