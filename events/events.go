package events

import (
	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/window"
)

type Events struct {
	WM *wm.WM
}

func New(wm *wm.WM) *Events {
	return &Events{WM: wm}
}

func (events *Events) MapRequest(win window.Window) error {
	client, err := client.New(win)
	if err != nil {
		return err
	}

	if err = events.WM.Manage(client); err != nil {
		return err
	}

	return nil
}

func (events *Events) Unmanage(id uint) error {
	if err := events.WM.Unmanage(id); err != nil {
		return err
	}
	return nil
}
