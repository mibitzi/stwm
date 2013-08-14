package xclient

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/test/events"
)

func TestWindowRunner(t *testing.T) {
	xu, _ := xgbutil.NewConnDisplay(":1")
	testId(t, xu)
	testManage(t, xu)
	testShowHide(t, xu)
	testGeom(t, xu)
}

func testId(t *testing.T, xu *xgbutil.XUtil) {
	win, _ := xwindow.Create(xu, xu.RootWin())
	client := New(xu, win.Id, events.New())

	assert.Equal(t, win.Id, client.Id(), "win.Id")
}

func testManage(t *testing.T, xu *xgbutil.XUtil) {
	win, _ := xwindow.Create(xu, xu.RootWin())
	client := New(xu, win.Id, events.New())

	assert.NoError(t, client.Manage(), "win.Manage")
}

func testShowHide(t *testing.T, xu *xgbutil.XUtil) {
	win, _ := xwindow.Create(xu, xu.RootWin())
	client := New(xu, win.Id, events.New())

	client.Hide()
	client.Show()
	assert.True(t, client.Visible(), "win.IsVisible")
	client.Hide()
	assert.False(t, client.Visible(), "win.IsVisible")
}

func testGeom(t *testing.T, xu *xgbutil.XUtil) {
	win, _ := xwindow.Create(xu, xu.RootWin())
	client := New(xu, win.Id, events.New())

	geom := rect.New(100, 200, 500, 600)
	client.SetGeom(geom)
	newGeom := client.Geom()
	assert.Equal(t, geom.X(), newGeom.X(), "x")
	assert.Equal(t, geom.Y(), newGeom.Y(), "y")
	assert.Equal(t, geom.Width(), newGeom.Width(), "width")
	assert.Equal(t, geom.Height(), newGeom.Height(), "height")
}
