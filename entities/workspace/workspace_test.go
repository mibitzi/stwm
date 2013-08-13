package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/layout/tiling"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/window"
)

func TestAddClient(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Show()

	client, _ := client.New(window.New())
	assert.NoError(t, ws.AddClient(client), "AddClient")
	assert.True(t, ws.HasClient(client.Id()), "ws.HasClient")
	assert.True(t, ws.tiling.HasClient(client.Id()), "ws.tiling.HasClient")
	assert.True(t, client.IsVisible(), "client.IsVisible")
}

func TestAddClientHidden(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Hide()

	client, _ := client.New(window.New())
	ws.AddClient(client)
	assert.False(t, client.IsVisible(), "client.IsVisible")
}

func TestShowHide(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Show()

	client0, _ := client.New(window.New())
	client1, _ := client.New(window.New())

	ws.AddClient(client0)
	ws.AddClient(client1)

	ws.Hide()
	assert.False(t, client0.IsVisible(), "client0.IsVisible")
	assert.False(t, client1.IsVisible(), "client1.IsVisible")

	ws.Show()
	assert.True(t, client0.IsVisible(), "client0.IsVisible")
	assert.True(t, client1.IsVisible(), "client1.IsVisible")
}

func TestRemoveClient(t *testing.T) {
	ws := New("ws", tiling.New(rect.New(0, 0, 100, 100)))
	ws.Show()

	client, _ := client.New(window.New())
	ws.AddClient(client)
	ws.RemoveClient(client)
	assert.False(t, ws.HasClient(client.Id()), "ws.HasClient")
	assert.False(t, ws.tiling.HasClient(client.Id()), "ws.tiling.HasClient")
}
