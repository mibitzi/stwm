package xwindow

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
	xwin, _ := xwindow.Create(xu, xu.RootWin())
	win := New(xu, xwin.Id, events.New())

	assert.Equal(t, xwin.Id, win.Id(), "win.Id")
}

func testManage(t *testing.T, xu *xgbutil.XUtil) {
	xwin, _ := xwindow.Create(xu, xu.RootWin())
	win := New(xu, xwin.Id, events.New())

	assert.NoError(t, win.Manage(), "win.Manage")
}

func testShowHide(t *testing.T, xu *xgbutil.XUtil) {
	xwin, _ := xwindow.Create(xu, xu.RootWin())
	win := New(xu, xwin.Id, events.New())

	win.Hide()
	win.Show()
	assert.True(t, win.IsVisible(), "win.IsVisible")
	win.Hide()
	assert.False(t, win.IsVisible(), "win.IsVisible")
}

func testGeom(t *testing.T, xu *xgbutil.XUtil) {
	xwin, _ := xwindow.Create(xu, xu.RootWin())
	win := New(xu, xwin.Id, events.New())

	geom := rect.New(100, 200, 500, 600)
	win.SetGeom(geom)
	newGeom := win.Geom()
	assert.Equal(t, geom.X(), newGeom.X(), "x")
	assert.Equal(t, geom.Y(), newGeom.Y(), "y")
	assert.Equal(t, geom.Width(), newGeom.Width(), "width")
	assert.Equal(t, geom.Height(), newGeom.Height(), "height")
}
