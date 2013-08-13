package window

import (
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/window"
)

type Window struct {
	id      uint
	geom    rect.Rect
	visible bool
}

var numWindows uint = 0

func New() window.Window {
	numWindows += 1

	return &Window{
		id:      numWindows,
		geom:    rect.New(0, 0, 0, 0),
		visible: false,
	}
}

func (win *Window) Manage() error {
	return nil
}

func (win *Window) Id() uint {
	return win.id
}

func (win *Window) Geom() rect.Rect {
	return win.geom
}

func (win *Window) SetGeom(geom rect.Rect) {
	win.geom = geom
}

func (win *Window) IsVisible() bool {
	return win.visible
}

func (win *Window) Show() {
	win.visible = true
}

func (win *Window) Hide() {
	win.visible = false
}
