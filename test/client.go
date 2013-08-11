package test

import (
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/mibitzi/stwm/client"
	"github.com/mibitzi/stwm/config"
)

func NewClient() *client.Client {
	win, _ := xwindow.Generate(X)
	config := config.New()
	return client.New(X, win.Id, config)
}
