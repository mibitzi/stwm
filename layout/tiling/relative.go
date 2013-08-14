package tiling

import (
	"fmt"

	"github.com/mibitzi/stwm/layout"
)

// Relative returns a another client which is in relative position to the given
// client. If there is no client in that direction the same client will be
// returned.
func (tiling *Tiling) Relative(client layout.Client,
	dir layout.Direction) (layout.Client, error) {
	switch dir {
	case layout.DIR_UP:
		return tiling.RelativeUp(client)
	case layout.DIR_DOWN:
		return tiling.RelativeDown(client)
	case layout.DIR_LEFT:
		return tiling.RelativeLeft(client)
	case layout.DIR_RIGHT:
		return tiling.RelativeRight(client)
	default:
		return nil, fmt.Errorf("tiling: unknown direction %v", dir)
	}
}

// RelativeUp returns the client above the given one.
func (tiling *Tiling) RelativeUp(client layout.Client) (layout.Client, error) {
	col, row, err := tiling.findClient(client.Id())
	if err != nil {
		return nil, err
	}

	if row == 0 {
		return client, nil
	} else {
		return tiling.columns[col].rows[row-1].client, nil
	}
}

// RelativeDown returns the client below the given one.
func (tiling *Tiling) RelativeDown(client layout.Client) (layout.Client,
	error) {
	col, row, err := tiling.findClient(client.Id())
	if err != nil {
		return nil, err
	}

	if row >= len(tiling.columns[col].rows)-1 {
		return client, nil
	} else {
		return tiling.columns[col].rows[row+1].client, nil
	}
}

// RelativeLeft returns the client left of the given one.
func (tiling *Tiling) RelativeLeft(client layout.Client) (layout.Client,
	error) {
	/*ecol, erow, err := tiling.findClient(client.Id())
	if err != nil {
		return nil, err
	}*/

	return nil, nil
}

// RelativeRight returns the client left of the given one.
func (tiling *Tiling) RelativeRight(client layout.Client) (layout.Client,
	error) {
	/*ecol, erow, err := tiling.findClient(client.Id())
	if err != nil {
		return nil, err
	}*/

	return nil, nil
}
