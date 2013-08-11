package wm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/test"
	"github.com/mibitzi/stwm/workspace"
)

func TestClientRunner(t *testing.T) {
	test.Run(t, func() {
		wm := createWM()
		wm.x = test.X
		wm.addWorkspace(workspace.New(wm.screenRect()))
		testManage(t, wm)
		testHasClient(t, wm)
		testFindClient(t, wm)
		testUnmanage(t, wm)
	})
}

func testManage(t *testing.T, wm *WM) {
	clients := len(wm.clients)
	assert.NoError(t, wm.manage(test.NewClient()), "Manage client")
	assert.Equal(t, clients+1, len(wm.clients), "Number of clients")
}

func testHasClient(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.manage(client)
	assert.True(t, wm.hasClient(client), "HasClient")
}

func testFindClient(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.manage(client)

	_, found, err := wm.findClient(client.Id())
	assert.NoError(t, err, "findClient")
	assert.Equal(t, client, found, "findClient")
}

func testUnmanage(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.manage(client)
	wm.unmanage(client.Id())
	assert.False(t, wm.hasClient(client), "HasClient")
	assert.Error(t, wm.unmanage(client.Id()))
}
