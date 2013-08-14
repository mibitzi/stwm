package xclient

import (
	"fmt"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/log"
	"github.com/mibitzi/stwm/rect"
)

type Events interface {
	Unmanage(uint)
}

type XClient struct {
	xu      *xgbutil.XUtil
	id      xproto.Window
	win     *xwindow.Window
	events  Events
	focused bool
}

func New(xu *xgbutil.XUtil, id xproto.Window, events Events) client.Client {
	return &XClient{
		xu:     xu,
		id:     id,
		win:    xwindow.New(xu, id),
		events: events,
	}
}

// Manage sets up all needed event handler for this window.
func (client *XClient) Manage() error {
	err := client.win.Listen(xproto.EventMaskEnterWindow,
		xproto.EventMaskPropertyChange)
	if err != nil {
		return fmt.Errorf("window: cannot listen to events: %s", err.Error())
	}

	xevent.UnmapNotifyFun(client.unmapNotify).Connect(client.xu, client.id)
	xevent.DestroyNotifyFun(client.destroyNotify).Connect(client.xu, client.id)

	return nil
}

// unmapNotify is the callback function for an UnmapNotifyEvent.
func (client *XClient) unmapNotify(xu *xgbutil.XUtil,
	ev xevent.UnmapNotifyEvent) {
	client.unmanage()
}

// destroyNotify is the callback function for a DestroyNotifyEvent.
func (client *XClient) destroyNotify(xu *xgbutil.XUtil,
	ev xevent.DestroyNotifyEvent) {
	client.unmanage()
}

// unmanage detaches all event handlers from this window.
func (client *XClient) unmanage() {
	client.win.Detach()
	client.events.Unmanage(uint(client.win.Id))
}

func (client *XClient) Id() uint {
	return uint(client.id)
}

func (client *XClient) Geom() rect.Rect {
	geom := client.win.Geom
	return rect.New(geom.X(), geom.Y(), geom.Width(), geom.Height())
}

func (client *XClient) SetGeom(geom rect.Rect) {
	client.win.MoveResize(geom.X(), geom.Y(), geom.Width(), geom.Height())
}

// Visible returns true if the window is mapped.
func (client *XClient) Visible() bool {
	if attr, err := xproto.GetWindowAttributes(client.xu.Conn(),
		client.win.Id).Reply(); err != nil {
		log.Warn(err)
		return false
	} else {
		return attr.MapState == xproto.MapStateViewable
	}
}

// Show maps the window.
func (client *XClient) Show() {
	client.win.Map()
}

// Hide unmaps the window.
func (client *XClient) Hide() {
	client.win.Unmap()
}

func (client *XClient) Focus() {
	client.win.Focus()
	client.focused = true
}

func (client *XClient) Unfocus() {
	client.focused = false
}

func (client *XClient) Focused() bool {
	return client.focused
}
