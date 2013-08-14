package tiling

import (
	"errors"

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

func New(geom rect.Rect) layout.Layout {
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
	row := &Row{client: client}
	if err := tiling.insertRow(0, row); err != nil {
		return err
	}
	tiling.apply()

	return nil
}

func (tiling *Tiling) InsertClient(client layout.Client, colIdx int,
	rowIdx int) error {
	return tiling.insertRowAfter(colIdx, rowIdx, &Row{client: client})
}

func (tiling *Tiling) RemoveClient(client layout.Client) error {
	if col, row, err := tiling.findClient(client.Id()); err != nil {
		return err
	} else {
		if err := tiling.removeRow(col, row); err != nil {
			return err
		}
		tiling.apply()
		return nil
	}
}

func (tiling *Tiling) removeRow(colIdx, rowIdx int) error {
	if colIdx < 0 || colIdx >= int(len(tiling.columns)) {
		return errors.New("tiling: column index out of bound")
	}

	col := tiling.columns[colIdx]
	col.rows = append(col.rows[:rowIdx], col.rows[rowIdx+1:]...)

	if len(tiling.columns) > 1 && len(col.rows) == 0 {
		tiling.columns = append(tiling.columns[:colIdx],
			tiling.columns[colIdx+1:]...)
		tiling.useFreeColumnSpace()
	} else {
		col.useFreeSpace(tiling.Geom.Height())
	}

	return nil
}

func (tiling *Tiling) insertRowAfter(colIdx, rowIdx int, row *Row) error {
	if colIdx < 0 || colIdx >= len(tiling.columns) {
		return errors.New("tiling: column index out of bound")
	}

	col := tiling.columns[colIdx]

	row.size = int(float64(tiling.Geom.Height()) / float64(len(col.rows)+1))

	col.makeSpace(row.size)

	if rowIdx < 0 || rowIdx >= len(col.rows)-1 {
		col.rows = append(col.rows, row)
	} else {
		idx := rowIdx + 1
		col.rows = append(col.rows, nil)
		copy(col.rows[idx+1:], col.rows[idx:])
		col.rows[idx] = row
	}

	return nil
}

func (tiling *Tiling) insertRow(colIdx int, row *Row) error {
	return tiling.insertRowAfter(colIdx, -1, row)
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
	return 0, 0, errors.New("tiling: client not found")
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

func (tiling *Tiling) makeColumnSpace(space int) {
	min := int(float64(space) / float64(len(tiling.columns)))

	for i, col := range tiling.columns {
		col.size -= min
		col.size -= space % (i + 1)
	}
}

func (col *Column) makeSpace(space int) {
	min := int(float64(space) / float64(len(col.rows)))

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

	freeSpace := total - usedSpace
	col.makeSpace(-freeSpace)
}

func (tiling *Tiling) useFreeColumnSpace() {
	usedSpace := 0
	for _, col := range tiling.columns {
		usedSpace += col.size
	}

	freeSpace := tiling.Geom.Width() - usedSpace
	tiling.makeColumnSpace(-freeSpace)
}
