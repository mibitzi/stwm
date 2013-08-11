package layout

import (
	"errors"
	"math"

	"github.com/BurntSushi/xgbutil/xrect"

	"github.com/mibitzi/stwm/client"
)

type tilingRow struct {
	size   int
	client *client.Client
}

type tilingColumn struct {
	size int
	rows []*tilingRow
}

type Tiling struct {
	Geom    xrect.Rect
	columns []*tilingColumn
}

func NewTiling(geom xrect.Rect) *Tiling {
	tiling := &Tiling{
		Geom:    geom,
		columns: make([]*tilingColumn, 1),
	}

	tiling.columns[0] = &tilingColumn{
		size: geom.Width(),
		rows: make([]*tilingRow, 0),
	}

	return tiling
}

func (tiling *Tiling) AddClient(client *client.Client) error {
	if err := tiling.insertClient(client, 0); err != nil {
		return err
	}

	tiling.apply()

	return nil
}

func (tiling *Tiling) RemoveClient(client *client.Client) error {
	if colIdx, rowIdx, err := tiling.findClient(client); err != nil {
		return err
	} else {
		col := tiling.columns[colIdx]
		col.rows[rowIdx] = col.rows[len(col.rows)-1]
		col.rows = col.rows[:len(col.rows)-1]
		col.useFreeSpace(tiling.Geom.Height())

		tiling.apply()
		return nil
	}
}

func (tiling *Tiling) HasClient(client *client.Client) bool {
	if _, _, err := tiling.findClient(client); err != nil {
		return false
	}
	return true
}

func (tiling *Tiling) findClient(client *client.Client) (int, int, error) {
	for colIdx, col := range tiling.columns {
		for rowIdx, row := range col.rows {
			if row.client == client {
				return colIdx, rowIdx, nil
			}
		}
	}
	return -1, -1, errors.New("Client not found")
}

func (tiling *Tiling) insertClient(client *client.Client, column int) error {
	if column < 0 || column >= len(tiling.columns) {
		return errors.New("Column index out of bound")
	}

	col := tiling.columns[column]

	row := &tilingRow{
		size:   tiling.Geom.Height() / (len(col.rows) + 1),
		client: client,
	}

	col.makeSpace(row.size)
	col.rows = append(col.rows, row)

	return nil
}

func (tiling *Tiling) apply() {
	left := 0
	for _, col := range tiling.columns {
		top := 0
		for _, row := range col.rows {
			row.client.MoveResize(left, top, col.size, row.size)
			top += row.size
		}
		left += col.size
	}
}

func (col *tilingColumn) makeSpace(space int) {
	min := int(math.Floor(float64(space) / float64(len(col.rows))))

	for i, row := range col.rows {
		row.size -= min
		row.size -= space % (i + 1)
	}
}

func (col *tilingColumn) useFreeSpace(total int) {
	usedSpace := 0
	for _, row := range col.rows {
		usedSpace += row.size
	}

	freeSpace := float64(total - usedSpace)

	for _, row := range col.rows {
		row.size += int(float64(row.size)/float64(usedSpace)*freeSpace + 0.5)
	}
}
