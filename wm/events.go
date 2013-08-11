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
		if err := wm.manage(client.New(wm.x, ev.Window)); err != nil {
			log.Print(err)
		}
	}).Connect(wm.x, wm.x.RootWin())
}

func (wm *WM) unmapNotifyFun(xu *xgbutil.XUtil, ev xevent.UnmapNotifyEvent) {
	log.Printf("UnmapNotifyEvent: %d\n", ev.Window)
	if err := wm.unmanage(ev.Window); err != nil {
		log.Print(err)
	}
}

func (wm *WM) destroyNotifyFun(xu *xgbutil.XUtil,
	ev xevent.DestroyNotifyEvent) {
	log.Printf("DestroyNotifyEvent: %d\n", ev.Window)
	if err := wm.unmanage(ev.Window); err != nil {
		log.Print(err)
	}
}
