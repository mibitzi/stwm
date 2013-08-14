package wm

import (
	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/entities/workspace"
)

type WM struct {
	Clients    []client.Client
	Focused    client.Client
	CurWs      *workspace.Workspace
	Workspaces []*workspace.Workspace
}

func New() (*WM, error) {
	wm := &WM{
		Clients:    make([]client.Client, 0),
		Workspaces: make([]*workspace.Workspace, 0),
	}

	return wm, nil
}

func (wm *WM) Destroy() {
}
