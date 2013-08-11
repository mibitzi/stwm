package wm

import (
	"fmt"
	"os/exec"
	"strings"
)

type CommandHandler map[string]func([]string) error

func NewCommandHandler() CommandHandler {
	handler := make(CommandHandler)
	handler["exec"] = handler.CmdExec

	return handler
}

func (handler CommandHandler) Execute(cmd string, args []string) error {
	if fn, ok := handler[cmd]; ok {
		if err := fn(args); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unknown command: %s", cmd)
	}

	return nil
}

// cmdExec executes the arguments given to it as shell commands.
func (handler CommandHandler) CmdExec(args []string) error {
	cmd := exec.Command("/bin/sh", "-c", strings.Join(args, ""))
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
