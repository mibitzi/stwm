package window

import (
	"github.com/mibitzi/stwm/rect"
)

type Window interface {
	Id() uint
	Manage() error
	SetGeom(rect.Rect)
	Geom() rect.Rect
	IsVisible() bool
	Show()
	Hide()
}
