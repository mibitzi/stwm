package client

import (
	"github.com/mibitzi/stwm/rect"
)

type Client interface {
	Id() uint
	Manage() error
	Geom() rect.Rect
	SetGeom(rect.Rect)
	Visible() bool
	Show()
	Hide()
	Focus()
	Unfocus()
	Focused() bool
}
