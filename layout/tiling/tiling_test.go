package tiling

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/window"
)

func TestHasClient(t *testing.T) {
	tiling := New(rect.New(0, 0, 1000, 1000))
	client := window.New()
	tiling.AddClient(client)
	assert.True(t, tiling.HasClient(client.Id()), "HasClient")
}

func TestAddClient(t *testing.T) {
	rect := rect.New(0, 0, 1000, 500)
	tiling := New(rect)

	client := window.New()
	tiling.AddClient(client)
	geom := client.Geom()

	assert.Equal(t, 0, geom.X(), "first window x")
	assert.Equal(t, 0, geom.Y(), "first window y")
	assert.Equal(t, rect.Width(), geom.Width(), "first window width")
	assert.Equal(t, rect.Height(), geom.Height(), "first window height")

	client = window.New()
	tiling.AddClient(client)
	geom = client.Geom()

	assert.Equal(t, 0, geom.X(), "second window x")
	assert.Equal(t, rect.Height()/2, geom.Y(), "second window y")
	assert.Equal(t, rect.Width(), geom.Width(), "second window width")
	assert.Equal(t, rect.Height()/2, geom.Height(), "second window height")

}

func TestRemoveClient(t *testing.T) {
	rect := rect.New(0, 0, 1000, 500)
	tiling := New(rect)

	client0 := window.New()
	client1 := window.New()
	tiling.AddClient(client0)
	tiling.AddClient(client1)

	assert.NoError(t, tiling.RemoveClient(client0), "RemoveClient")
	assert.False(t, tiling.HasClient(client0.Id()), "HasClient")

	geom := client1.Geom()
	assert.Equal(t, 0, geom.X(), "window x")
	assert.Equal(t, 0, geom.Y(), "window y")
	assert.Equal(t, rect.Width(), geom.Width(), "window width")
	assert.Equal(t, rect.Height(), geom.Height(), "window height")
}
