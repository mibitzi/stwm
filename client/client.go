package client

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type Client struct {
	X          *xgbutil.XUtil
	Win        *xwindow.Window
	Floating   bool
	Fullscreen bool
	Urgent     bool
}

func New(xu *xgbutil.XUtil, wid xproto.Window) *Client {
	client := &Client{
		X:   xu,
		Win: xwindow.New(xu, wid),
	}

	client.Win.Listen(xproto.EventMaskEnterWindow |
		xproto.EventMaskFocusChange |
		xproto.EventMaskPropertyChange)

	return client
}

func (client *Client) Id() xproto.Window {
	return client.Win.Id
}
