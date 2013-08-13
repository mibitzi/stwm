package xgb

import (
	"fmt"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/mibitzi/stwm/rect"
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

func (xgb *Xgb) ScreenRect() rect.Rect {
	return rect.New(0, 0, int(xgb.X.Screen().WidthInPixels),
		int(xgb.X.Screen().HeightInPixels))
}
