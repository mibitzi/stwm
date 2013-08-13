package commands

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/wm"
)

func newTestCommands() *Commands {
	wm, _ := wm.New()
	return New(wm)
}

func testCmdExec(t *testing.T) {
	cmd := newTestCommands()

	file := "/tmp/stwm.test"
	os.Remove(file)

	cmd.Execute(fmt.Sprintf("exec touch %s", file))
	exec.Command("sync").Run()

	_, err := os.Stat(file)
	assert.NoError(t, err)
	os.Remove(file)
}
