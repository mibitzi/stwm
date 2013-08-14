package events

import (
	"github.com/mibitzi/stwm/entities/client"
)

type Events struct{}

func New() *Events {
	return &Events{}
}

func (ev *Events) MapRequest(client client.Client) {
}

func (ev *Events) Unmanage(id uint) {
}
