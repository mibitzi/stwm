package events

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/commands"
	"github.com/mibitzi/stwm/test/window"
)

func TestMapRequest(t *testing.T) {
	wm, _ := wm.New()
	ws := workspace.New("1", tiling.New(rect.New(0, 0, 0, 0)))
	wm.AddWorkspace(ws)

	ev := New(wm, commands.New())

	win := window.New()
	assert.NoError(t, ev.MapRequest(win), "MapRequest")
	assert.True(t, wm.HasClient(win.Id()), "wm.HasClient")
}

func TestUnmanage(t *testing.T) {
	wm, _ := wm.New()
	ws := workspace.New("1", tiling.New(rect.New(0, 0, 0, 0)))
	wm.AddWorkspace(ws)

	ev := New(wm, commands.New())
	win := window.New()
	ev.MapRequest(win)
	ev.Unmanage(win.Id())
	assert.False(t, wm.HasClient(win.Id()), "wm.HasClient")
}
