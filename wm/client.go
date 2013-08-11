package wm

import (
	"errors"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/client"
)

func (wm *WM) Manage(client *client.Client) error {
	if wm.CurWs == nil {
		return errors.New("No current workspace available")
	}

	if wm.HasClient(client) {
		return errors.New("Client already managed")
	}

	wm.Clients = append(wm.Clients, client)
	xevent.UnmapNotifyFun(wm.unmapNotifyFun).Connect(wm.X, client.Win.Id)
	xevent.DestroyNotifyFun(wm.destroyNotifyFun).Connect(wm.X, client.Win.Id)

	wm.CurWs.AddClient(client)
	client.Win.Map()

	return nil
}

func (wm *WM) Unmanage(id xproto.Window) error {
	if idx, client, err := wm.findClient(id); err != nil {
		return err
	} else {
		wm.Clients[idx] = wm.Clients[len(wm.Clients)-1]
		wm.Clients = wm.Clients[:len(wm.Clients)-1]

		for _, ws := range wm.Wspaces {
			if ws.HasClient(client) {
				if err := ws.RemoveClient(client); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func (wm *WM) HasClient(client *client.Client) bool {
	if _, _, err := wm.findClient(client.Win.Id); err != nil {
		return false
	}
	return true
}

func (wm *WM) findClient(id xproto.Window) (int, *client.Client, error) {
	for i, c := range wm.Clients {
		if c.Id() == id {
			return i, c, nil
		}
	}
	return -1, nil, errors.New("Client not found")
}
