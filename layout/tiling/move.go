package tiling

import (
	"fmt"

	"github.com/mibitzi/stwm/layout"
)

// Move moves a client into a direction.
func (tiling *Tiling) Move(client layout.Client, dir layout.Direction) error {
	switch dir {
	case layout.DIR_UP:
		return tiling.MoveUp(client)
	case layout.DIR_DOWN:
		return tiling.MoveDown(client)
	case layout.DIR_LEFT:
		return tiling.MoveLeft(client)
	case layout.DIR_RIGHT:
		return tiling.MoveRight(client)
	default:
		return fmt.Errorf("tiling: unknown direction %v", dir)
	}
}

// MoveLeft moves a client one column to the left, by removing it from its
// column and inserting it at the and of the column left of it.
// If the client was not found an error will be returned.
// If the client is already alone in the leftmost column nothing happens.
func (tiling *Tiling) MoveLeft(client layout.Client) error {
	return tiling.moveLeftRight(client, -1)
}

// MoveRight moves a client one column to the right, by removing it from its
// column and inserting it at the and of the column left of it.
// If the client was not found an error will be returned.
// If the client is already alone in the rightmost column nothing happens.
func (tiling *Tiling) MoveRight(client layout.Client) error {
	return tiling.moveLeftRight(client, 1)
}

// moveLeftRight moves a client to the left if dir < 0, or to the right if
// dir > 0
func (tiling *Tiling) moveLeftRight(client layout.Client, dir int) error {
	colIdx, rowIdx, err := tiling.findClient(client.Id())
	if err != nil {
		return err
	}

	column := tiling.columns[colIdx]
	row := column.rows[rowIdx]

	outermost := ((dir < 0 && colIdx == 0) ||
		(dir > 0 && colIdx == len(tiling.columns)-1))

	if outermost && len(column.rows) == 1 {
		// There is no point in moving this row
		return nil
	} else if outermost {
		// We are in the outermost column. Therefore we need to create a new
		// column for ourselfs.

		tiling.removeRow(colIdx, rowIdx)
		row.size = tiling.Geom.Height()

		// Prepare a reasonable amount of space
		size := int(float64(tiling.Geom.Width()) /
			float64(len(tiling.columns)+1))
		tiling.makeColumnSpace(size)

		// Insert a new column
		col := &Column{size: size, rows: make([]*Row, 1)}
		col.rows[0] = row

		if dir < 0 {
			tiling.columns = append(tiling.columns, nil)
			copy(tiling.columns[1:], tiling.columns[:])
			tiling.columns[0] = col
		} else {
			tiling.columns = append(tiling.columns, col)
		}
	} else {
		// In all other cases we just shift the row, by removing it from its
		// column and appending it to the new column.

		tiling.removeRow(colIdx, rowIdx)
		tiling.insertRow(colIdx+dir, row)
	}

	tiling.apply()

	return nil
}

// MoveDown moves a client down one row.
// If its already in the bottommost row, nothing happens.
func (tiling *Tiling) MoveDown(client layout.Client) error {
	colIdx, rowIdx, err := tiling.findClient(client.Id())
	if err != nil {
		return err
	}

	col := tiling.columns[colIdx]
	if rowIdx == len(col.rows)-1 {
		return nil
	}

	row := col.rows[rowIdx]
	col.rows = append(col.rows[:rowIdx], col.rows[rowIdx+1:]...)

	idx := rowIdx + 1
	col.rows = append(col.rows, nil)
	copy(col.rows[idx+1:], col.rows[idx:])
	col.rows[idx] = row

	tiling.apply()

	return nil
}

// MoveUp moves a client up one row.
// If its already in the topmost row, nothing happens.
func (tiling *Tiling) MoveUp(client layout.Client) error {
	col, row, err := tiling.findClient(client.Id())
	if err != nil {
		return err
	}

	if row == 0 {
		return nil
	}

	return tiling.MoveDown(tiling.columns[col].rows[row-1].client)
}
