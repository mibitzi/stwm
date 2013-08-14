package workspace

import (
	"errors"

	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/layout"
)

type Workspace struct {
	id      string
	clients []client.Client
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
func (ws *Workspace) AddClient(client client.Client) error {
	if ws.HasClient(client.Id()) {
		return errors.New("workspace: already added this client")
	}

	ws.clients = append(ws.clients, client)

	ws.tiling.AddClient(client)

	if ws.Visible() {
		client.Show()
	}

	return nil
}

// RemoveClients removes a client from this workspace.
func (ws *Workspace) RemoveClient(client client.Client) error {
	idx, err := ws.findClient(client.Id())
	if err != nil {
		return err
	}

	// No need to keep the order
	ws.clients[idx] = ws.clients[len(ws.clients)-1]
	ws.clients = ws.clients[:len(ws.clients)-1]

	if err = ws.tiling.RemoveClient(client); err != nil {
		return err
	}

	return nil
}

// HasClient checks if this workspace has a client with the given id.
func (ws *Workspace) HasClient(id uint) bool {
	_, err := ws.findClient(id)
	return err == nil
}

// findClient returns the index of a client.
func (ws *Workspace) findClient(id uint) (int, error) {
	for idx, c := range ws.clients {
		if c.Id() == id {
			return idx, nil
		}
	}
	return -1, errors.New("workspace: client not found")
}

// IsVisible returns true if this workspace is currently visible.
func (ws *Workspace) Visible() bool {
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

// MoveClient moves a client into a direction.
func (ws *Workspace) MoveClient(client client.Client, dir string) error {
	return ws.tiling.Move(client, layout.Direction(dir))
}
