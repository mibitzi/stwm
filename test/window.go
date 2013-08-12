package test

type Window struct {
	id uint
}

var numWindows uint = 0

func NewWindow() *Window {
	numWindows += 1
	return &Window{id: numWindows}
}

func (win *Window) Id() uint {
	return win.id
}
