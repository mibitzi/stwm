package main

import (
	"log"

	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/events"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/xgb"
)

func main() {
	var err error

	wm, err := wm.New()
	if err != nil {
		log.Fatal(err)
	}
	defer wm.Destroy()

	xgb, err := xgb.New(events.New(wm))
	if err != nil {
		log.Fatal(err)
	}
	defer xgb.Destroy()

	wm.AddWorkspace(workspace.New("1", tiling.New(rect.New(0, 0, 1024, 768))))

	xgb.EnterMain()
}
