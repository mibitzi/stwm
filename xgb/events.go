package xgb

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/log"
	"github.com/mibitzi/stwm/window"
	"github.com/mibitzi/stwm/window/xwindow"
)

type Events interface {
	MapRequest(window.Window) error
	xwindow.Events
	Command(string) error
}

func (xgb *Xgb) setupEvents() {
	xevent.MapRequestFun(func(xu *xgbutil.XUtil, ev xevent.MapRequestEvent) {
		win := xwindow.New(xgb.X, ev.Window, xgb.Events)
		if err := xgb.Events.MapRequest(win); err != nil {
			log.Error(err)
		}
	}).Connect(xgb.X, xgb.X.RootWin())

	xgb.setupKeys()
}

func (xgb *Xgb) setupKeys() error {
	keybind.Initialize(xgb.X)

	for key, _ := range xgb.Config.Keybinds {
		log.Debug("handling key", key)
		if err := keybind.KeyPressFun(xgb.keyPressFun).Connect(xgb.X,
			xgb.X.RootWin(), key, true); err != nil {
			log.Error(err)
		}
	}

	return nil
}

func (xgb *Xgb) keyPressFun(xu *xgbutil.XUtil, e xevent.KeyPressEvent) {
	modStr := keybind.ModifierString(e.State)
	keyStr := keybind.LookupString(xgb.X, e.State, e.Detail)

	var keys string

	if len(modStr) > 0 {
		keys = fmt.Sprintf("%s-%s", modStr, keyStr)
	} else {
		keys = keyStr
	}

	keys = strings.ToLower(keys)

	if command, ok := xgb.Config.Keybinds[keys]; ok {
		if err := xgb.Events.Command(command); err != nil {
			log.Error(err)
		}
	}
}
