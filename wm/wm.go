package wm

import (
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/client"
	"github.com/mibitzi/stwm/config"
	"github.com/mibitzi/stwm/workspace"
)

type WM struct {
	x       *xgbutil.XUtil
	clients []*client.Client
	curWs   *workspace.Workspace
	wspaces []*workspace.Workspace
	config  *config.Config
}

func SetupWM() *WM {
	wm := createWM()

	wm.config = config.New()

	wm.connect()
	wm.setupRoot()
	wm.setupEvents()
	wm.setupKeys()
	wm.setupWorkspaces()

	return wm
}

func createWM() *WM {
	return &WM{
		x:       nil,
		clients: make([]*client.Client, 0),
		curWs:   nil,
		wspaces: make([]*workspace.Workspace, 0),
	}

}

func (wm *WM) Start() {
	xevent.Main(wm.x)
}

func (wm *WM) Destroy() {
	wm.x.Conn().Close()
}

func (wm *WM) connect() {
	x, err := xgbutil.NewConnDisplay(":1")
	if err != nil {
		log.Fatal(err)
	}

	wm.x = x
}
