package xwindow

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/mibitzi/stwm/rect"
)

type Events interface {
	Unmanage(uint) error
}

type XWindow struct {
	xu     *xgbutil.XUtil
	id     xproto.Window
	xwin   *xwindow.Window
	events Events
}

func New(xu *xgbutil.XUtil, id xproto.Window, events Events) *XWindow {
	return &XWindow{
		xu:     xu,
		id:     id,
		xwin:   xwindow.New(xu, id),
		events: events,
	}
}

// Manage sets up all needed event handler for this window.
func (win *XWindow) Manage() error {
	err := win.xwin.Listen(xproto.EventMaskEnterWindow,
		xproto.EventMaskPropertyChange)
	if err != nil {
		return fmt.Errorf("window: cannot listen to events: %s", err.Error())
	}

	xevent.UnmapNotifyFun(win.unmapNotify).Connect(win.xu, win.id)
	xevent.DestroyNotifyFun(win.destroyNotify).Connect(win.xu, win.id)

	return nil
}

// unmapNotify is the callback function for an UnmapNotifyEvent.
func (win *XWindow) unmapNotify(xu *xgbutil.XUtil,
	ev xevent.UnmapNotifyEvent) {
	win.unmanage()
}

// destroyNotify is the callback function for a DestroyNotifyEvent.
func (win *XWindow) destroyNotify(xu *xgbutil.XUtil,
	ev xevent.DestroyNotifyEvent) {
	win.unmanage()
}

// unmanage detaches all event handlers from this window.
func (win *XWindow) unmanage() {
	win.xwin.Detach()
	if err := win.events.Unmanage(uint(win.xwin.Id)); err != nil {
		log.Print(err)
	}
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
