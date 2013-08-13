package events

import (
	"github.com/mibitzi/stwm/window"
)

type Events struct{}

func New() *Events {
	return &Events{}
}

func (ev *Events) MapRequest(win window.Window) {
}

func (ev *Events) Unmanage(id uint) error {
	return nil
}
