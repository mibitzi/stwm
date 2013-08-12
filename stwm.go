package stwm

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/mibitzi/stwm/entities/wm"
)

func main() {
	wm := wm.New()
	defer wm.Destroy()

	connect()

	setupEvents()
}

func connect() {
	if xu, err := xgbutil.NewConnDisplay(":1"); err != nil {
		return nil, err
	}
}
