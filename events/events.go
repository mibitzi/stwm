package events

import (
	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/window"
)

type Events struct {
	WM  *wm.WM
	Cmd CommandHandler
}

type CommandHandler interface {
	Execute(string) error
}

func New(wm *wm.WM, cmd CommandHandler) *Events {
	return &Events{
		WM:  wm,
		Cmd: cmd,
	}
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

func (events *Events) Command(str string) error {
	return events.Cmd.Execute(str)
}
