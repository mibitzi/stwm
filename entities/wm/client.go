package wm

import (
	"errors"
	"github.com/mibitzi/stwm/entities/client"
)

// Manage adds a new managed client to this wm.
func (wm *WM) Manage(client *client.Client) error {
	if wm.HasClient(client.Id()) {
		return errors.New("manage: already managing a client with this id")
	}

	if len(wm.Workspaces) == 0 {
		return errors.New("manage: no workspaces available to manage client")
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
