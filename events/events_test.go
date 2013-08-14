package events

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/client"
	"github.com/mibitzi/stwm/test/commands"
)

func TestMapRequest(t *testing.T) {
	wm, _ := wm.New()
	ws := workspace.New("1", tiling.New(rect.New(0, 0, 0, 0)))
	wm.AddWorkspace(ws)

	ev := New(wm, commands.New())

	client := client.New()
	ev.MapRequest(client)
	assert.True(t, wm.HasClient(client.Id()), "wm.HasClient")
}

func TestUnmanage(t *testing.T) {
	wm, _ := wm.New()
	ws := workspace.New("1", tiling.New(rect.New(0, 0, 0, 0)))
	wm.AddWorkspace(ws)

	ev := New(wm, commands.New())
	client := client.New()
	ev.MapRequest(client)
	ev.Unmanage(client.Id())
	assert.False(t, wm.HasClient(client.Id()), "wm.HasClient")
}
