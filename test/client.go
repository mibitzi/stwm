package test

import (
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/mibitzi/stwm/client"
)

func NewClient() *client.Client {
	win, _ := xwindow.Generate(X)
	return client.New(X, win.Id)
}
