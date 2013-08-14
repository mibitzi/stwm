package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/client"
)

func TestAddClient(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Show()

	client := client.New()
	assert.NoError(t, ws.AddClient(client), "AddClient")
	assert.True(t, ws.HasClient(client.Id()), "ws.HasClient")
	assert.True(t, ws.tiling.HasClient(client.Id()), "ws.tiling.HasClient")
	assert.True(t, client.Visible(), "client.IsVisible")
}

func TestAddClientHidden(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Hide()

	client := client.New()
	ws.AddClient(client)
	assert.False(t, client.Visible(), "client.IsVisible")
}

func TestShowHide(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Show()

	client0 := client.New()
	client1 := client.New()

	ws.AddClient(client0)
	ws.AddClient(client1)

	ws.Hide()
	assert.False(t, client0.Visible(), "client0.IsVisible")
	assert.False(t, client1.Visible(), "client1.IsVisible")

	ws.Show()
	assert.True(t, client0.Visible(), "client0.IsVisible")
	assert.True(t, client1.Visible(), "client1.IsVisible")
}

func TestRemoveClient(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Show()

	client := client.New()
	ws.AddClient(client)
	ws.RemoveClient(client)
	assert.False(t, ws.HasClient(client.Id()), "ws.HasClient")
	assert.False(t, ws.tiling.HasClient(client.Id()), "ws.tiling.HasClient")
}
