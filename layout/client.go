package layout

import (
	"github.com/mibitzi/stwm/rect"
)

type Client interface {
	Id() uint
	Geom() rect.Rect
	SetGeom(rect.Rect)
}
