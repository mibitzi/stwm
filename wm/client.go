package wm

import (
	"errors"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/client"
)

func (wm *WM) manage(client *client.Client) error {
	if wm.curWs == nil {
		return errors.New("No current workspace available")
	}

	if wm.hasClient(client) {
		return errors.New("Client already managed")
	}

	wm.clients = append(wm.clients, client)
	xevent.UnmapNotifyFun(wm.unmapNotifyFun).Connect(wm.x, client.Win.Id)
	xevent.DestroyNotifyFun(wm.destroyNotifyFun).Connect(wm.x, client.Win.Id)

	wm.curWs.AddClient(client)
	client.Win.Map()

	return nil
}

func (wm *WM) unmanage(id xproto.Window) error {
	if idx, client, err := wm.findClient(id); err != nil {
		return err
	} else {
		wm.clients[idx] = wm.clients[len(wm.clients)-1]
		wm.clients = wm.clients[:len(wm.clients)-1]

		for _, ws := range wm.wspaces {
			if ws.HasClient(client) {
				if err := ws.RemoveClient(client); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func (wm *WM) hasClient(client *client.Client) bool {
	if _, _, err := wm.findClient(client.Win.Id); err != nil {
		return false
	}
	return true
}

func (wm *WM) findClient(id xproto.Window) (int, *client.Client, error) {
	for i, c := range wm.clients {
		if c.Id() == id {
			return i, c, nil
		}
	}
	return -1, nil, errors.New("Client not found")
}
