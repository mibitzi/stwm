package client

import (
	"github.com/mibitzi/stwm/rect"
	"github.com/mibitzi/stwm/window"
)

type Client struct {
	win window.Window
}

func New(win window.Window) (*Client, error) {
	client := &Client{
		win: win,
	}

	return client, nil
}

func (client *Client) Manage() error {
	return client.win.Manage()
}

func (client *Client) Id() uint {
	return uint(client.win.Id())
}

func (client *Client) Geom() rect.Rect {
	return client.win.Geom()
}

func (client *Client) SetGeom(rect rect.Rect) {
	client.win.SetGeom(rect)
}

func (client *Client) IsVisible() bool {
	return client.win.IsVisible()
}

func (client *Client) Show() {
	client.win.Show()
}

func (client *Client) Hide() {
	client.win.Hide()
}
