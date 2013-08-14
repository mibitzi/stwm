package layout

import ()

type Direction string

const (
	DIR_UP    = "up"
	DIR_DOWN  = "down"
	DIR_LEFT  = "left"
	DIR_RIGHT = "right"
)

type Layout interface {
	AddClient(Client) error
	InsertClient(Client, int, int) error
	RemoveClient(Client) error
	HasClient(uint) bool
	Move(Client, Direction) error
	Relative(Client, Direction) (Client, error)
}
