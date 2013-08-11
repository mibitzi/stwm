package wm

import (
	"github.com/BurntSushi/xgbutil/xrect"
)

func (wm *WM) screenRect() xrect.Rect {
	return xrect.New(0, 0, int(wm.x.Screen().WidthInPixels),
		int(wm.x.Screen().HeightInPixels))
}
