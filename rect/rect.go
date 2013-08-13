package rect

import (
	"github.com/BurntSushi/xgbutil/xrect"
)

type Rect interface {
	X() int
	Y() int
	Width() int
	Height() int
	SetX(int)
	SetY(int)
	SetWidth(int)
	SetHeight(int)
}

type rectImp struct {
	*xrect.XRect
}

func New(x, y, width, height int) *rectImp {
	return &rectImp{xrect.New(x, y, width, height)}
}

func (rect *rectImp) SetX(x int) {
	rect.XSet(x)
}

func (rect *rectImp) SetY(y int) {
	rect.YSet(y)
}

func (rect *rectImp) SetWidth(width int) {
	rect.WidthSet(width)
}

func (rect *rectImp) SetHeight(height int) {
	rect.HeightSet(height)
}
