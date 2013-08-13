package window

import (
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/mibitzi/stwm/rect"
)

type Window interface {
	Id() uint
	SetGeom(rect.Rect)
	Geom() rect.Rect
	IsVisible() bool
	Show()
	Hide()
}

type XWindow struct {
	xu   *xgbutil.XUtil
	id   xproto.Window
	xwin *xwindow.Window
}

func New(xu *xgbutil.XUtil, id xproto.Window) Window {
	return &XWindow{xu: xu, id: id, xwin: xwindow.New(xu, id)}
}

func (win *XWindow) Id() uint {
	return uint(win.id)
}

func (win *XWindow) Geom() rect.Rect {
	geom := win.xwin.Geom
	return rect.New(geom.X(), geom.Y(), geom.Width(), geom.Height())
}

func (win *XWindow) SetGeom(geom rect.Rect) {
	win.xwin.MoveResize(geom.X(), geom.Y(), geom.Width(), geom.Height())
}

// IsVisible returns true if the window is mapped.
func (win *XWindow) IsVisible() bool {
	if attr, err := xproto.GetWindowAttributes(win.xu.Conn(),
		win.xwin.Id).Reply(); err != nil {
		log.Print(err)
		return false
	} else {
		return attr.MapState == xproto.MapStateViewable
	}
}

// Show maps the window.
func (win *XWindow) Show() {
	win.xwin.Map()
}

// Hide unmaps the window.
func (win *XWindow) Hide() {
	win.xwin.Unmap()
}
