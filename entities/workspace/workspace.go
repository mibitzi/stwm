package workspace

import (
	"errors"

	"github.com/mibitzi/stwm/entities"
)

type Workspace struct {
	id      string
	clients []entities.Client
}

func New(id string) *Workspace {
	ws := &Workspace{id: id}
	return ws
}

func (ws *Workspace) Id() string {
	return ws.id
}

func (ws *Workspace) AddClient(client entities.Client) error {
	if ws.HasClient(client.Id()) {
		return errors.New("workspace: already added this client")
	}

	ws.clients = append(ws.clients, client)

	return nil
}

func (ws *Workspace) HasClient(id uint) bool {
	for _, c := range ws.clients {
		if c.Id() == id {
			return true
		}
	}
	return false
}
