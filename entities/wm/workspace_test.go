package wm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
)

func TestAddWorkspace(t *testing.T) {
	wm, _ := New()
	ws := workspace.New("1", tiling.New(rect.New(0, 0, 0, 0)))
	assert.NoError(t, wm.AddWorkspace(ws), "wm.AddWorkspace")
	assert.True(t, wm.HasWorkspace("1"), "wm.HasWorkspace")
	assert.Equal(t, ws, wm.CurWs, "wm.CurWs")
	assert.True(t, wm.CurWs.IsVisible(), "wm.CurWs.IsVisible")
}
