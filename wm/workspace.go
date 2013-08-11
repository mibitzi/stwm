package wm

import (
	"github.com/mibitzi/stwm/workspace"
)

func (wm *WM) setupWorkspaces() {
	wm.addWorkspace(workspace.New(wm.screenRect()))
}

func (wm *WM) addWorkspace(ws *workspace.Workspace) {
	wm.wspaces = append(wm.wspaces, ws)

	if len(wm.wspaces) == 1 {
		wm.curWs = ws
	}
}
