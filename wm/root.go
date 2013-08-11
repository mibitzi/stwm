package wm

import (
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xwindow"
)

func (wm *WM) setupRoot() {
	eventMask := xproto.EventMaskStructureNotify |
		xproto.EventMaskSubstructureNotify |
		xproto.EventMaskSubstructureRedirect |
		xproto.EventMaskEnterWindow

	if err := xwindow.New(wm.X, wm.X.RootWin()).Listen(eventMask); err != nil {
		log.Fatalf("Could not listen to root window events: %s", err)
	}
}
