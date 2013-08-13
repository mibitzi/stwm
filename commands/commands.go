package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mibitzi/stwm/entities/wm"
)

type HandlerMap map[string]func(args string) error

type Commands struct {
	WM      *wm.WM
	Handler HandlerMap
}

func New(wm *wm.WM) *Commands {
	cmd := &Commands{
		WM:      wm,
		Handler: make(HandlerMap),
	}

	cmd.Handler["exec"] = cmd.CmdExec

	return cmd
}

func (cmd *Commands) Execute(str string) error {
	parts := strings.SplitN(str, " ", 2)

	if handler, ok := cmd.Handler[parts[0]]; ok {
		return handler(parts[1])
	} else {
		return fmt.Errorf("commands: unknown command %s", parts[0])
	}
}

func (cmd *Commands) CmdExec(args string) error {
	if err := exec.Command("/bin/sh", "-c", args).Start(); err != nil {
		return err
	}
	return nil
}
