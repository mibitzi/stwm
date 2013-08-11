package wm

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

func (wm *WM) setupKeys() error {
	keybind.Initialize(wm.x)

	for key, _ := range wm.config.Keybinds {
		log.Printf("handling key %s\n", key)
		if err := keybind.KeyPressFun(wm.keyPressFun).Connect(wm.x,
			wm.x.RootWin(), key, true); err != nil {
			log.Print(err)
		}
	}

	return nil
}

func (wm *WM) keyPressFun(xu *xgbutil.XUtil, e xevent.KeyPressEvent) {
	modStr := keybind.ModifierString(e.State)
	keyStr := keybind.LookupString(wm.x, e.State, e.Detail)

	var keys string

	if len(modStr) > 0 {
		keys = fmt.Sprintf("%s-%s", modStr, keyStr)
	} else {
		keys = keyStr
	}

	keys = strings.ToLower(keys)

	log.Printf("got key %s\n", keys)

	if command, ok := wm.config.Keybinds[keys]; ok {
		if err := wm.executeCommand(command.Cmd, command.Args); err != nil {
			log.Print(err)
		}
	}
}
