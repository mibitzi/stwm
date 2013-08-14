package client

import (
	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/rect"
)

type Client struct {
	id      uint
	geom    rect.Rect
	visible bool
	focused bool
}

var numClients uint = 0

func New() client.Client {
	numClients += 1

	return &Client{
		id:      numClients,
		geom:    rect.New(0, 0, 0, 0),
		visible: false,
	}
}

func (client *Client) Manage() error {
	return nil
}

func (client *Client) Id() uint {
	return client.id
}

func (client *Client) Geom() rect.Rect {
	return client.geom
}

func (client *Client) SetGeom(geom rect.Rect) {
	client.geom = geom
}

func (client *Client) Visible() bool {
	return client.visible
}

func (client *Client) Show() {
	client.visible = true
}

func (client *Client) Hide() {
	client.visible = false
}

func (client *Client) Focus() {
	client.focused = true
}

func (client *Client) Unfocus() {
	client.focused = false
}

func (client *Client) Focused() bool {
	return client.focused
}
