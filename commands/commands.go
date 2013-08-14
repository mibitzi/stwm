package commands

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mibitzi/stwm/entities/wm"
	"github.com/mibitzi/stwm/log"
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
	cmd.Handler["move"] = cmd.CmdMove

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

func (cmd *Commands) CmdMove(args string) error {
	focused := cmd.WM.Focused
	if focused == nil {
		log.Debug("command: called move without focused window")
		return nil
	}

	if ws, err := cmd.WM.ClientWorkspace(cmd.WM.Focused.Id()); err != nil {
		return err
	} else {
		return ws.MoveClient(focused, strings.ToLower(args))
	}
}
