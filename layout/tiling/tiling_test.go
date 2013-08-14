package tiling

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/layout"
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/client"
)

func assertPos(t *testing.T, tiling *Tiling, client layout.Client,
	col int, row int, msg string) {

	cidx, ridx, _ := tiling.findClient(client.Id())

	assert.Equal(t, col, cidx, fmt.Sprintf("%s: column", msg))
	assert.Equal(t, row, ridx, fmt.Sprintf("%s: row", msg))
}

func assertGeom(t *testing.T, tiling *Tiling, client layout.Client,
	rect rect.Rect, msg string) {
	geom := client.Geom()
	assert.Equal(t, rect.X(), geom.X(), fmt.Sprintf("%s: x", msg))
	assert.Equal(t, rect.Y(), geom.Y(), fmt.Sprintf("%s: y", msg))
	assert.Equal(t, rect.Width(), geom.Width(), fmt.Sprintf("%s: width", msg))
	assert.Equal(t, rect.Height(), geom.Height(),
		fmt.Sprintf("%s: height", msg))
}

func prepareTestClients(cnt int) (*Tiling, []layout.Client, rect.Rect) {
	rect := rect.New(0, 0, 1000, 600)
	tiling := New(rect)

	clients := make([]layout.Client, cnt)

	for i, _ := range clients {
		clients[i] = client.New()
		tiling.AddClient(clients[i])
	}

	return tiling.(*Tiling), clients, rect
}

func TestHasClient(t *testing.T) {
	tiling := New(rect.New(0, 0, 1000, 1000))
	client := client.New()
	assert.NoError(t, tiling.AddClient(client))
	assert.True(t, tiling.HasClient(client.Id()), "HasClient")
}

func TestAddClient(t *testing.T) {
	rec := rect.New(0, 0, 1000, 500)
	tiling := New(rec).(*Tiling)

	cl := client.New()
	tiling.AddClient(cl)

	assertGeom(t, tiling, cl, rect.New(0, 0, rec.Width(), rec.Height()),
		"first")

	cl = client.New()
	tiling.AddClient(cl)

	assertGeom(t, tiling, cl, rect.New(0, rec.Height()/2, rec.Width(),
		rec.Height()/2), "second")
}

func TestRemoveClient(t *testing.T) {
	rec := rect.New(0, 0, 1000, 500)
	tiling := New(rec).(*Tiling)

	client0 := client.New()
	client1 := client.New()
	tiling.AddClient(client0)
	tiling.AddClient(client1)

	assert.NoError(t, tiling.RemoveClient(client0), "RemoveClient")
	assert.False(t, tiling.HasClient(client0.Id()), "HasClient")

	assertGeom(t, tiling, client1, rect.New(0, 0, rec.Width(), rec.Height()),
		"remaining window")
}

func TestInsertClient(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)

	client := client.New()
	tiling.InsertClient(client, 0, 1)
	assertPos(t, tiling, clients[0], 0, 0, "1st client")
	assertPos(t, tiling, clients[1], 0, 1, "2nd client")
	assertPos(t, tiling, client, 0, 2, "3rd client")
	assertPos(t, tiling, clients[2], 0, 3, "4th client")
}
