package main

import (
	"github.com/mibitzi/stwm/wm"
)

func main() {
	wm := wm.SetupWM()
	defer wm.Destroy()
	wm.Start()
}
