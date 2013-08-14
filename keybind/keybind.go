package keybinds

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/mibitzi/stwm/log"
)

type Events interface {
	Command(string)
}

type Key struct {
	Mods uint16
	Code xproto.Keycode
	Cmd  string
}

type Keybinds struct {
	X      *xgbutil.XUtil
	Keys   []*Key
	Events Events
}

// New creates a new Keybinds instance. It receives a map which contains
// keystrings with commands.
func New(xu *xgbutil.XUtil, events Events, keys map[string]string) *Keybinds {
	keybind.Initialize(xu)

	kb := &Keybinds{
		X:      xu,
		Events: events,
		Keys:   make([]*Key, 0, len(keys)),
	}

	for keystr, cmd := range keys {
		mods, kcs, err := keybind.ParseString(xu, keystr)
		if err != nil {
			log.Warn(err)
			continue
		}

		key := &Key{
			Mods: mods,
			Code: kcs[0],
			Cmd:  cmd,
		}

		kb.Keys = append(kb.Keys, key)

		if err := keybind.KeyPressFun(kb.keyPressFun).Connect(xu,
			xu.RootWin(), keystr, true); err != nil {
			log.Warn(err)
		}
	}

	return kb
}

func (kb *Keybinds) Destroy() {
}

func (kb *Keybinds) keyPressFun(xu *xgbutil.XUtil, ev xevent.KeyPressEvent) {
	for _, key := range kb.Keys {
		if key.Mods == ev.State && key.Code == ev.Detail {
			kb.Events.Command(key.Cmd)
			return
		}
	}

	log.Warn("keybind: received key without keybind")
}
