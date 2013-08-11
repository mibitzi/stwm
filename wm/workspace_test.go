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
		wm.X = test.X
		testAddWorkspace(t, wm)
	})
}

func testAddWorkspace(t *testing.T, wm *WM) {
	wm.Wspaces = make([]*workspace.Workspace, 0)
	wm.CurWs = nil

	ws := workspace.New(wm.ScreenRect())
	wm.AddWorkspace(ws)

	assert.Equal(t, 1, len(wm.Wspaces), "Number of workspaces")
	assert.Equal(t, wm.CurWs, ws, "Current workspace")
}
