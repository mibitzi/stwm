package test

import (
	"testing"

	"github.com/BurntSushi/xgbutil"
)

var (
	X *xgbutil.XUtil
)

func Run(t *testing.T, fn func()) {
	if X == nil {
		xu, err := xgbutil.NewConnDisplay(":1")
		if err != nil {
			t.Fatal(err)
		}

		X = xu
	}

	fn()
}
