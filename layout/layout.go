package layout

import ()

type Layout interface {
	AddClient(Client) error
	RemoveClient(Client) error
	HasClient(uint) bool
}
