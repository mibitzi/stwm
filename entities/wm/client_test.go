package wm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/entities/workspace"
	"github.com/mibitzi/stwm/test"
)

func TestManage(t *testing.T) {
	wm, _ := New()
	wm.AddWorkspace(workspace.New("1"))

	win := test.NewWindow()
	client, _ := client.New(win)
	wm.Manage(client)
	assert.True(t, wm.HasClient(client.Id()), "new: WM has client")
	assert.True(t, wm.CurWs.HasClient(client.Id()), "new: CurWs HasClient")
}
