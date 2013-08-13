package xgb

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
)

type Xgb struct {
	X      *xgbutil.XUtil
	events Events
}

func New(events Events) (*Xgb, error) {
	xgb := &Xgb{events: events}

	if err := xgb.connect(); err != nil {
		return nil, err
	}

	if err := xgb.setupRoot(); err != nil {
		return nil, err
	}

	xgb.setupEvents()

	return xgb, nil
}

func (xgb *Xgb) connect() error {
	var err error
	if xgb.X, err = xgbutil.NewConnDisplay(":1"); err != nil {
		return err
	}

	return nil
}

func (xgb *Xgb) EnterMain() {
	xevent.Main(xgb.X)
}

func (xgb *Xgb) Destroy() {
	xgb.X.Conn().Close()
}
