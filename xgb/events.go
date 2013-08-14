package xgb

import (
	//"fmt"
	//"strings"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/entities/client"
	//"github.com/mibitzi/stwm/log"
	"github.com/mibitzi/stwm/xclient"
)

type Events interface {
	MapRequest(client.Client)
	xclient.Events
	//Command(string)
}

func (xgb *Xgb) setupEvents() {
	xevent.MapRequestFun(func(xu *xgbutil.XUtil, ev xevent.MapRequestEvent) {
		client := xclient.New(xgb.X, ev.Window, xgb.Events)
		xgb.Events.MapRequest(client)
	}).Connect(xgb.X, xgb.X.RootWin())

	//xgb.setupKeys()
}

/*
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
		xgb.Events.Command(command)
	} else {
		log.Warn("xgb: received key without keybind:", keys)
	}
}*/
