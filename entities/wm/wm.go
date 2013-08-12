package wm

import (
	"github.com/mibitzi/stwm/entities"
)

type WM struct {
	Clients    []entities.Client
	CurWs      entities.Workspace
	Workspaces []entities.Workspace
}

func New() (*WM, error) {
	wm := &WM{
		Clients:    make([]entities.Client, 0),
		Workspaces: make([]entities.Workspace, 0),
	}

	return wm, nil
}

func (wm *WM) Destroy() {
}
