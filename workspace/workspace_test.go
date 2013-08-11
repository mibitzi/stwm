package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BurntSushi/xgbutil/xrect"

	"github.com/mibitzi/stwm/test"
)

func TestWorkspaceRunner(t *testing.T) {
	test.Run(t, func() {
		testHasClient(t)
		testRemoveClient(t)
	})
}

func testHasClient(t *testing.T) {
	ws := New(xrect.New(0, 0, 0, 0))
	client := test.NewClient()
	assert.NoError(t, ws.AddClient(client))
	assert.True(t, ws.HasClient(client), "HasClient")
}

func testRemoveClient(t *testing.T) {
	ws := New(xrect.New(0, 0, 0, 0))
	client := test.NewClient()
	ws.AddClient(client)
	assert.NoError(t, ws.RemoveClient(client))
	assert.False(t, ws.HasClient(client), "HasClient")
}
