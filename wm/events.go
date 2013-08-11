package wm

import (
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/client"
)

func (wm *WM) setupEvents() {
	xevent.MapRequestFun(func(xu *xgbutil.XUtil, ev xevent.MapRequestEvent) {
		log.Printf("MapRequestEvent: %d\n", ev.Window)
		if err := wm.Manage(client.New(wm.X, ev.Window,
			wm.Config)); err != nil {
			log.Print(err)
		}
	}).Connect(wm.X, wm.X.RootWin())
}

func (wm *WM) unmapNotifyFun(xu *xgbutil.XUtil, ev xevent.UnmapNotifyEvent) {
	log.Printf("UnmapNotifyEvent: %d\n", ev.Window)
	if err := wm.Unmanage(ev.Window); err != nil {
		log.Print(err)
	}
}

func (wm *WM) destroyNotifyFun(xu *xgbutil.XUtil,
	ev xevent.DestroyNotifyEvent) {
	log.Printf("DestroyNotifyEvent: %d\n", ev.Window)
	if err := wm.Unmanage(ev.Window); err != nil {
		log.Print(err)
	}
}
