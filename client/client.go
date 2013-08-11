package client

import (
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xrect"
	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/mibitzi/stwm/config"
)

type Client struct {
	X          *xgbutil.XUtil
	Win        *xwindow.Window
	Floating   bool
	Fullscreen bool
	Urgent     bool
	Focused    bool
	Geom       xrect.Rect

	config *config.Config
}

func New(xu *xgbutil.XUtil, wid xproto.Window, config *config.Config) *Client {
	client := &Client{
		X:      xu,
		Win:    xwindow.New(xu, wid),
		config: config,
	}

	client.Geom = client.Win.Geom

	client.Win.Listen(xproto.EventMaskEnterWindow |
		xproto.EventMaskFocusChange |
		xproto.EventMaskPropertyChange)

	if borderWidth, err := config.IntVar("borderWidth"); err != nil {
		log.Print(err)
	} else {
		xproto.ConfigureWindow(client.X.Conn(), client.Win.Id,
			uint16(xproto.ConfigWindowBorderWidth),
			[]uint32{uint32(borderWidth)})
	}

	return client
}

func (client *Client) Id() xproto.Window {
	return client.Win.Id
}

// Focus sets the input focus to this client and changes the windows appearance
// accordingly.
func (client *Client) Focus() {
	client.Focused = true
	client.Win.Focus()
	client.setBorderColor()
}

func (client *Client) Unfocus() {
	client.Focused = false
	client.setBorderColor()
}

func (client *Client) setBorderColor() {
	var varName string

	if client.Focused {
		varName = "activeWindowColor"
	} else {
		varName = "inactiveWindowColor"
	}

	if color, err := client.config.IntVar(varName); err != nil {
		log.Print(err)
	} else {
		log.Printf("set borderpixel: %v", color)
		client.Win.Change(xproto.CwBorderPixel, uint32(color))
	}
}

func (client *Client) MoveResize(x, y, width, height int) {
	client.Geom = xrect.New(x, y, width, height)

	borderWidth, _ := client.config.IntVar("borderWidth")
	client.Win.MoveResize(x, y, width-2*borderWidth, height-2*borderWidth)
}
