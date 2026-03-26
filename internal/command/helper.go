package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/matteo-gildone/sigil/internal/cli"
	"github.com/matteo-gildone/sigil/internal/store"
	"github.com/matteo-gildone/sigil/internal/xdg"
)

func withStore(project string, fn func(*store.Store, []byte) error) error {
	passphrase, err := prompt("passphrase:")
	if err != nil {
		return fmt.Errorf("failed to read passphrase: %w", err)
	}
	// zero passphrase after use; best-effort due to GO's GC
	defer func() {
		for i := range passphrase {
			passphrase[i] = 0
		}
	}()

	path, err := xdg.ProjectPath(project)
	if err != nil {
		return fmt.Errorf("failed to read project path: %w", err)
	}

	s, err := store.Load(path, passphrase)
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	return fn(s, passphrase)
}

func prompt(label string) ([]byte, error) {
	p, err := cli.PromptPassphrase(label, int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", strings.TrimSuffix(label, ":"), err)
	}
	return p, nil
}

func clipboardCommand(tool string) *exec.Cmd {
	args := strings.Fields(tool)
	return exec.Command(args[0], args[1:]...)
}

func copyToClipboard(tool, value, key string, clearAfter int) error {
	cmd := clipboardCommand(tool)
	cmd.Stdin = strings.NewReader(value)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy to clipboard: %w", err)
	}
	fmt.Fprintf(os.Stdout, "%q copied to clipboard\n", key)
	if clearAfter > 0 {
		fmt.Fprintf(os.Stdout, "clearing clipboard in %d seconds\n", clearAfter)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(clearAfter) * time.Second)
			cmd := clipboardCommand(tool)
			cmd.Stdin = strings.NewReader("")
			if err := cmd.Run(); err != nil {
				fmt.Fprintln(os.Stderr, "failed to clear clipboard")
			}
		}()

		wg.Wait()
		fmt.Fprintln(os.Stdout, "clipboard cleared")
	}
	return nil
}
