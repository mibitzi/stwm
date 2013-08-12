package entities

type Workspace interface {
	Id() string
	AddClient(Client) error
	HasClient(uint) bool
}
