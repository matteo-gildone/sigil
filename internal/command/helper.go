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

type multiFlag []string

func (f *multiFlag) String() string {
	return strings.Join(*f, ", ")
}

func (f *multiFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func loadStore(project string) (*store.Store, []byte, error) {
	passphrase, err := prompt("passphrase:")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read passphrase: %w", err)
	}

	path, err := xdg.ProjectPath(project)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read project path: %w", err)
	}

	s, err := store.Load(path, string(passphrase))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load store: %w", err)
	}

	return s, passphrase, nil
}

func prompt(label string) ([]byte, error) {
	p, err := cli.PromptPassphrase(label, int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", strings.TrimSuffix(label, ":"), err)
	}
	return p, nil
}

func copyToClipboard(value, key string, clearAfter int) error {
	cmd := exec.Command("pbcopy")
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
			cmd := exec.Command("pbcopy")
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
