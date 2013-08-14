package events

import (
	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/log"
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

func (events *Events) MapRequest(client client.Client) {
	if err := events.WM.Manage(client); err != nil {
		log.Error(err)
	}
}

func (events *Events) Unmanage(id uint) {
	if err := events.WM.Unmanage(id); err != nil {
		log.Error(err)
	}
}

func (events *Events) Command(str string) {
	log.Debug("events: got command", str)
	if err := events.Cmd.Execute(str); err != nil {
		log.Error(err)
	}
}
