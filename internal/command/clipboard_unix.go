//go:build !windows

package command

import (
	"os/exec"
	"runtime"
)

func clipboardCommand() *exec.Cmd {
	if runtime.GOOS == "darwin" {
		return exec.Command("pbcopy")
	}

	return exec.Command("xclip", "-selection", "clipboard")
}
