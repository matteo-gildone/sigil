package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/matteo-gildone/gostyl"
)

type ClipboardWriter interface {
	Write(data []byte) error
}

type execWriter struct {
	tool string
}

func (w *execWriter) Write(data []byte) error {
	cmd := clipboardCommand(w.tool)
	cmd.Stdin = bytes.NewReader(data)
	return cmd.Run()
}

func clipboardCommand(tool string) *exec.Cmd {
	args := strings.Fields(tool)
	return exec.Command(args[0], args[1:]...)
}

func copyToClipboard(writer ClipboardWriter, key, value string, clearAfter int) error {
	err := writer.Write([]byte(value))
	if err != nil {
		return fmt.Errorf("failed to copy to clipboard: %w", err)
	}

	fmt.Fprint(os.Stdout, gostyl.Successf("%q copied to clipboard\n", key))
	if clearAfter > 0 {
		fmt.Fprint(os.Stdout, gostyl.Infof("clearing clipboard in %d seconds\n", clearAfter))
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(clearAfter) * time.Second)
			err := writer.Write([]byte(""))
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to clear clipboard")
			}
		}()

		wg.Wait()
		fmt.Fprintln(os.Stdout, gostyl.Successln("clipboard cleared"))
	}
	return nil
}
