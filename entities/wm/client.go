package wm

import (
	"errors"
	"github.com/mibitzi/stwm/entities"
)

// Manage adds a new managed client to this wm.
func (wm *WM) Manage(client entities.Client) error {
	if wm.HasClient(client.Id()) {
		return errors.New("client: already created a client with this id")
	}

	wm.Clients = append(wm.Clients, client)

	wm.CurWs.AddClient(client)

	return nil
}

// HasClient checks if a client with the given id was already added to this wm.
func (wm *WM) HasClient(id uint) bool {
	for _, client := range wm.Clients {
		if client.Id() == id {
			return true
		}
	}
	return false
}
