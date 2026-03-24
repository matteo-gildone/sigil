//go:build windows

package command

import "os/exec"

func clipboardCommand() *exec.Cmd {
	return exec.Command("clip")
}
