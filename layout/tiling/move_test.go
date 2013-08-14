package tiling

import (
	"testing"

	"github.com/mibitzi/stwm/rect"
)

func TestMoveLeftSimple(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	tiling.MoveLeft(clients[0])
	assertPos(t, tiling, clients[0], 0, 0, "left client")
	assertPos(t, tiling, clients[1], 1, 0, "1st right client")
	assertPos(t, tiling, clients[2], 1, 1, "2nd right client")
}

func TestMoveLeftOutermost(t *testing.T) {
	tiling, clients, _ := prepareTestClients(2)
	tiling.MoveLeft(clients[0])
	tiling.MoveLeft(clients[0])
	assertPos(t, tiling, clients[0], 0, 0, "left client")
	assertPos(t, tiling, clients[1], 1, 0, "1st right client")
}

func TestMoveLeftShift(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	tiling.MoveLeft(clients[0])
	tiling.MoveLeft(clients[1])
	assertPos(t, tiling, clients[0], 0, 0, "1st left client")
	assertPos(t, tiling, clients[1], 0, 1, "2nd left client")
	assertPos(t, tiling, clients[2], 1, 0, "1st right client")
}

func TestMoveRightOutermost(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	tiling.MoveRight(clients[2])
	tiling.MoveRight(clients[2])
	assertPos(t, tiling, clients[0], 0, 0, "1st left client")
	assertPos(t, tiling, clients[1], 0, 1, "2nd left client")
	assertPos(t, tiling, clients[2], 1, 0, "1st right client")
}

func TestMoveDown(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	tiling.MoveDown(clients[1])
	assertPos(t, tiling, clients[0], 0, 0, "1st client")
	assertPos(t, tiling, clients[1], 0, 2, "2nd client")
	assertPos(t, tiling, clients[2], 0, 1, "3rd client")
}

func TestMoveBottommost(t *testing.T) {
	tiling, clients, _ := prepareTestClients(2)
	tiling.MoveDown(clients[1])
	assertPos(t, tiling, clients[0], 0, 0, "1st client")
	assertPos(t, tiling, clients[1], 0, 1, "2nd client")
}

func TestMoveUp(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	tiling.MoveUp(clients[1])
	assertPos(t, tiling, clients[0], 0, 1, "1st client")
	assertPos(t, tiling, clients[1], 0, 0, "2nd client")
	assertPos(t, tiling, clients[2], 0, 2, "3rd client")
}

func TestMoveUpTopmost(t *testing.T) {
	tiling, clients, _ := prepareTestClients(2)
	tiling.MoveUp(clients[0])
	assertPos(t, tiling, clients[0], 0, 0, "1st client")
	assertPos(t, tiling, clients[1], 0, 1, "2nd client")
}

func TestMoveLeftRightSize(t *testing.T) {
	tiling, clients, rec := prepareTestClients(3)
	tiling.MoveRight(clients[1])

	assertGeom(t, tiling, clients[1], rect.New(rec.Width()/2, 0, rec.Width()/2,
		rec.Height()), "moved right")

	tiling.MoveLeft(clients[1])
	assertGeom(t, tiling, clients[1], rect.New(0, rec.Height()/3*2,
		rec.Width(), rec.Height()/3), "moved left")
}
