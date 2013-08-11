package wm

import (
	"fmt"
	"os/exec"
	"strings"
)

func (wm *WM) executeCommand(cmd string, args []string) error {

	handlers := make(map[string]func([]string) error)
	handlers["exec"] = wm.cmdExec

	if fn, ok := handlers[cmd]; ok {
		if err := fn(args); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unknown command: %s", cmd)
	}

	return nil
}

// cmdExec executes the arguments given to it as shell commands.
func (wm *WM) cmdExec(args []string) error {
	cmd := exec.Command("/bin/sh", "-c", strings.Join(args, ""))
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
