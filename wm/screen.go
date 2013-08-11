package wm

import (
	"github.com/BurntSushi/xgbutil/xrect"
)

func (wm *WM) ScreenRect() xrect.Rect {
	return xrect.New(0, 0, int(wm.X.Screen().WidthInPixels),
		int(wm.X.Screen().HeightInPixels))
}
