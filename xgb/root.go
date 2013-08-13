package xgb

import (
	"fmt"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xwindow"
)

func (xgb *Xgb) setupRoot() error {
	eventMask := xproto.EventMaskStructureNotify |
		xproto.EventMaskSubstructureNotify |
		xproto.EventMaskSubstructureRedirect |
		xproto.EventMaskEnterWindow

	win := xwindow.New(xgb.X, xgb.X.RootWin())
	if err := win.Listen(eventMask); err != nil {
		return fmt.Errorf("Could not listen to root window events: %s",
			err.Error())
	}
	return nil
}
