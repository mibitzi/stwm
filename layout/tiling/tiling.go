package tiling

import (
	"errors"
	"math"

	"github.com/mibitzi/stwm/layout"
	"github.com/mibitzi/stwm/rect"
)

type Row struct {
	size   int
	client layout.Client
}

type Column struct {
	size int
	rows []*Row
}

type Tiling struct {
	Geom    rect.Rect
	columns []*Column
}

func New(geom rect.Rect) *Tiling {
	tiling := &Tiling{
		Geom:    geom,
		columns: make([]*Column, 1),
	}

	tiling.columns[0] = &Column{
		size: geom.Width(),
		rows: make([]*Row, 0),
	}

	return tiling
}

func (tiling *Tiling) AddClient(client layout.Client) error {
	if err := tiling.insertClient(client, 0); err != nil {
		return err
	}

	tiling.apply()

	return nil
}

func (tiling *Tiling) RemoveClient(client layout.Client) error {
	if colIdx, rowIdx, err := tiling.findClient(client.Id()); err != nil {
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

func (tiling *Tiling) HasClient(id uint) bool {
	if _, _, err := tiling.findClient(id); err != nil {
		return false
	}
	return true
}

func (tiling *Tiling) findClient(id uint) (int, int, error) {
	for colIdx, col := range tiling.columns {
		for rowIdx, row := range col.rows {
			if row.client.Id() == id {
				return colIdx, rowIdx, nil
			}
		}
	}
	return -1, -1, errors.New("tiling: client not found")
}

func (tiling *Tiling) insertClient(client layout.Client, column int) error {
	if column < 0 || column >= len(tiling.columns) {
		return errors.New("tiling: column index out of bound")
	}

	col := tiling.columns[column]

	row := &Row{
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
			row.client.SetGeom(rect.New(left, top, col.size, row.size))
			top += row.size
		}
		left += col.size
	}
}

func (col *Column) makeSpace(space int) {
	min := int(math.Floor(float64(space) / float64(len(col.rows))))

	for i, row := range col.rows {
		row.size -= min
		row.size -= space % (i + 1)
	}
}

func (col *Column) useFreeSpace(total int) {
	usedSpace := 0
	for _, row := range col.rows {
		usedSpace += row.size
	}

	freeSpace := float64(total - usedSpace)

	for _, row := range col.rows {
		row.size += int(float64(row.size)/float64(usedSpace)*freeSpace + 0.5)
	}
}
