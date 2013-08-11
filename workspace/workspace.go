package workspace

import (
	"github.com/BurntSushi/xgbutil/xrect"

	"github.com/mibitzi/stwm/client"
	"github.com/mibitzi/stwm/layout"
)

type Workspace struct {
	Rect   xrect.Rect
	Tiling *layout.Tiling
}

func New(rect xrect.Rect) *Workspace {
	ws := &Workspace{
		Rect:   rect,
		Tiling: layout.NewTiling(rect),
	}

	return ws
}

func (ws *Workspace) AddClient(client *client.Client) error {
	ws.Tiling.AddClient(client)
	return nil
}

func (ws *Workspace) RemoveClient(client *client.Client) error {
	return ws.Tiling.RemoveClient(client)
}

func (ws *Workspace) HasClient(client *client.Client) bool {
	return ws.Tiling.HasClient(client)
}
