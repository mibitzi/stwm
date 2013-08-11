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
	X       *xgbutil.XUtil
	Clients []*client.Client
	Focused *client.Client
	CurWs   *workspace.Workspace
	Wspaces []*workspace.Workspace
	Config  *config.Config
	Cmd     CommandHandler
}

func SetupWM() *WM {
	wm := createWM()

	wm.Config = config.New()
	wm.Cmd = NewCommandHandler()

	wm.connect()
	wm.setupRoot()
	wm.setupEvents()
	wm.setupKeys()
	wm.setupWorkspaces()

	return wm
}

func createWM() *WM {
	return &WM{
		Clients: make([]*client.Client, 0),
		Wspaces: make([]*workspace.Workspace, 0),
	}

}

func (wm *WM) Start() {
	xevent.Main(wm.X)
}

func (wm *WM) Destroy() {
	wm.X.Conn().Close()
}

func (wm *WM) connect() {
	x, err := xgbutil.NewConnDisplay(":1")
	if err != nil {
		log.Fatal(err)
	}

	wm.X = x
}
