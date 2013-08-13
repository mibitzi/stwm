package test

import (
	"github.com/mibitzi/stwm/rect"
)

type Window struct {
	id      uint
	geom    rect.Rect
	visible bool
}

var numWindows uint = 0

func NewWindow() *Window {
	numWindows += 1

	return &Window{
		id:      numWindows,
		geom:    rect.New(0, 0, 0, 0),
		visible: false,
	}
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
