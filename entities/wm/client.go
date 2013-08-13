package wm

import (
	"errors"

	"github.com/mibitzi/stwm/entities/client"
)

// Manage adds a new managed client to this wm.
func (wm *WM) Manage(client *client.Client) error {
	if wm.HasClient(client.Id()) {
		return errors.New("wm: manage: already managing a client with this id")
	}

	if len(wm.Workspaces) == 0 {
		return errors.New("wm: manage: no workspaces available")
	}

	if err := wm.CurWs.AddClient(client); err != nil {
		return err
	}

	if err := client.Manage(); err != nil {
		return err
	}

	wm.Clients = append(wm.Clients, client)

	return nil
}

// Unmanage removes a client from this wm.
func (wm *WM) Unmanage(id uint) error {
	idx, err := wm.findClient(id)
	if err != nil {
		return err
	}

	for _, ws := range wm.Workspaces {
		ws.RemoveClient(wm.Clients[idx])
	}

	wm.Clients[idx] = wm.Clients[len(wm.Clients)-1]
	wm.Clients = wm.Clients[:len(wm.Clients)-1]

	return nil
}

// HasClient checks if a client with the given id was already added to this wm.
func (wm *WM) HasClient(id uint) bool {
	_, err := wm.findClient(id)
	return err == nil
}

// findClient finds a client by its id.
func (wm *WM) findClient(id uint) (int, error) {
	for idx, client := range wm.Clients {
		if client.Id() == id {
			return idx, nil
		}
	}
	return -1, errors.New("wm: client not found")
}
