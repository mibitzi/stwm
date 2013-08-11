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
		wm.X = test.X
		wm.AddWorkspace(workspace.New(wm.ScreenRect()))

		testManage(t, wm)
		testHasClient(t, wm)
		testFindClient(t, wm)
		testUnmanage(t, wm)
		testFocus(t, wm)
	})
}

func testManage(t *testing.T, wm *WM) {
	numClients := len(wm.Clients)
	client := test.NewClient()
	assert.NoError(t, wm.Manage(client), "Manage client")
	assert.Equal(t, numClients+1, len(wm.Clients), "Number of clients")
	assert.Equal(t, wm.Focused, client, "Focused client")
}

func testHasClient(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.Manage(client)
	assert.True(t, wm.HasClient(client), "HasClient")
}

func testFindClient(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.Manage(client)

	_, found, err := wm.findClient(client.Id())
	assert.NoError(t, err, "findClient")
	assert.Equal(t, client, found, "findClient")
}

func testUnmanage(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.Manage(client)
	wm.Unmanage(client.Id())
	assert.False(t, wm.HasClient(client), "HasClient")
	assert.Error(t, wm.Unmanage(client.Id()))
}

func testFocus(t *testing.T, wm *WM) {
	client := test.NewClient()
	wm.Manage(client)
	wm.Manage(test.NewClient())
	wm.Focus(client)
	assert.Equal(t, wm.Focused, client, "Focused client")
	assert.True(t, client.Focused, "Focused flag")
}
