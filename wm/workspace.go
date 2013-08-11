package wm

import (
	"github.com/mibitzi/stwm/workspace"
)

func (wm *WM) setupWorkspaces() {
	wm.AddWorkspace(workspace.New(wm.ScreenRect()))
}

func (wm *WM) AddWorkspace(ws *workspace.Workspace) {
	wm.Wspaces = append(wm.Wspaces, ws)

	if len(wm.Wspaces) == 1 {
		wm.CurWs = ws
	}
}
