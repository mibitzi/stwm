package client

type Client struct {
	win Window
}

func New(win Window) (*Client, error) {
	client := &Client{
		win: win,
	}

	return client, nil
}

func (client *Client) Id() uint {
	return uint(client.win.Id())
}
