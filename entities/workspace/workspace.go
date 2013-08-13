package workspace

import (
	"errors"

	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/layout"
)

type Workspace struct {
	id      string
	clients []*client.Client
	tiling  layout.Layout
	visible bool
}

func New(id string, tiling layout.Layout) *Workspace {
	ws := &Workspace{id: id, tiling: tiling}
	return ws
}

// Id returns the id of this workspace.
func (ws *Workspace) Id() string {
	return ws.id
}

// AddClient adds a new client to this workspace and assigns it to a suitable
// layout.
func (ws *Workspace) AddClient(client *client.Client) error {
	if ws.HasClient(client.Id()) {
		return errors.New("workspace: already added this client")
	}

	ws.clients = append(ws.clients, client)

	ws.tiling.AddClient(client)

	if ws.IsVisible() {
		client.Show()
	}

	return nil
}

// HasClient checks if this workspace has a client with the given id.
func (ws *Workspace) HasClient(id uint) bool {
	for _, c := range ws.clients {
		if c.Id() == id {
			return true
		}
	}
	return false
}

// IsVisible returns true if this workspace is currently visible.
func (ws *Workspace) IsVisible() bool {
	return ws.visible
}

// Show shows this workspace and all clients in it.
func (ws *Workspace) Show() {
	ws.visible = true

	for _, c := range ws.clients {
		c.Show()
	}
}

// Hide hides this workspace and all clients in it.
func (ws *Workspace) Hide() {
	ws.visible = false

	for _, c := range ws.clients {
		c.Hide()
	}
}
