package layout

import (
	"testing"

	"github.com/BurntSushi/xgbutil/xrect"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/test"
)

func TestTilingRunner(t *testing.T) {
	test.Run(t, func() {
		testHasClient(t)
		testAddClient(t)
		testRemoveClient(t)
	})
}

func testHasClient(t *testing.T) {
	tiling := NewTiling(xrect.New(0, 0, 1000, 1000))
	client := test.NewClient()
	tiling.AddClient(client)
	assert.True(t, tiling.HasClient(client), "HasClient")
}

func testAddClient(t *testing.T) {
	rect := xrect.New(0, 0, 1000, 500)
	tiling := NewTiling(rect)

	client := test.NewClient()
	tiling.AddClient(client)
	geom := client.Geom

	assert.Equal(t, 0, geom.X(), "first window x")
	assert.Equal(t, 0, geom.Y(), "first window y")
	assert.Equal(t, rect.Width(), geom.Width(), "first window width")
	assert.Equal(t, rect.Height(), geom.Height(), "first window height")

	client = test.NewClient()
	tiling.AddClient(client)
	geom = client.Geom

	assert.Equal(t, 0, geom.X(), "second window x")
	assert.Equal(t, rect.Height()/2, geom.Y(), "second window y")
	assert.Equal(t, rect.Width(), geom.Width(), "second window width")
	assert.Equal(t, rect.Height()/2, geom.Height(), "second window height")

}

func testRemoveClient(t *testing.T) {
	rect := xrect.New(0, 0, 1000, 500)
	tiling := NewTiling(rect)

	client0 := test.NewClient()
	client1 := test.NewClient()
	tiling.AddClient(client0)
	tiling.AddClient(client1)

	assert.NoError(t, tiling.RemoveClient(client0), "RemoveClient")
	assert.False(t, tiling.HasClient(client0), "HasClient")

	geom := client1.Geom
	assert.Equal(t, 0, geom.X(), "window x")
	assert.Equal(t, 0, geom.Y(), "window y")
	assert.Equal(t, rect.Width(), geom.Width(), "window width")
	assert.Equal(t, rect.Height(), geom.Height(), "window height")
}
