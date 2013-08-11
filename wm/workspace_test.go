package wm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/test"
	"github.com/mibitzi/stwm/workspace"
)

func TestWorkspaceRunner(t *testing.T) {
	test.Run(t, func() {
		wm := createWM()
		wm.x = test.X
		testAddWorkspace(t, wm)
	})
}

func testAddWorkspace(t *testing.T, wm *WM) {
	wm.wspaces = make([]*workspace.Workspace, 0)
	wm.curWs = nil

	ws := workspace.New(wm.screenRect())
	wm.addWorkspace(ws)

	assert.Equal(t, 1, len(wm.wspaces), "Number of workspaces")
	assert.Equal(t, wm.curWs, ws, "Current workspace")
}
