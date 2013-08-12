package wm

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/test"
	"github.com/mibitzi/stwm/workspace"
)

func TestCommandsRunner(t *testing.T) {
	test.Run(t, func() {
		wm := createWM()
		wm.X = test.X
		wm.AddWorkspace(workspace.New(wm.ScreenRect()))
		wm.Cmd = NewCommandHandler()

		testCmdExec(t, wm)
	})
}

func testCmdExec(t *testing.T, wm *WM) {
	file := "/tmp/stwm.test"
	os.Remove(file)

	wm.Cmd.Execute("exec", []string{fmt.Sprintf("touch %s", file)})
	exec.Command("sync").Run()

	_, err := os.Stat(file)
	assert.NoError(t, err)
	os.Remove(file)
}
