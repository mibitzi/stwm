package wm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/client"
)

func TestManage(t *testing.T) {
	wm, _ := New()
	wm.AddWorkspace(workspace.New("1", tiling.New(rect.New(0, 0, 0, 0))))

	client := client.New()
	wm.Manage(client)
	assert.True(t, wm.HasClient(client.Id()), "new: WM has client")
	assert.True(t, wm.CurWs.HasClient(client.Id()), "new: CurWs HasClient")
	assert.Equal(t, wm.Focused, client, "wm.Focused")
	assert.True(t, client.Focused(), "client.Focused")
}

func TestClientWorkspace(t *testing.T) {
	wm, _ := New()
	ws := workspace.New("1", tiling.New(rect.New(0, 0, 0, 0)))
	wm.AddWorkspace(ws)

	client := client.New()
	wm.Manage(client)
	found, err := wm.ClientWorkspace(client.Id())
	assert.NoError(t, err, "wm.ClientWorkspace")
	assert.Equal(t, ws, found, "client workspace")
}
