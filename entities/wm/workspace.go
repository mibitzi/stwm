package wm

import (
	"errors"

	"github.com/mibitzi/stwm/entities/workspace"
)

func (wm *WM) AddWorkspace(ws *workspace.Workspace) error {
	if wm.HasWorkspace(ws.Id()) {
		return errors.New("wm: already added this workspace")
	}

	wm.Workspaces = append(wm.Workspaces, ws)

	if wm.CurWs == nil {
		wm.CurWs = ws
		wm.CurWs.Show()
	}

	return nil
}

func (wm *WM) HasWorkspace(id string) bool {
	for _, ws := range wm.Workspaces {
		if ws.Id() == id {
			return true
		}
	}
	return false
}
