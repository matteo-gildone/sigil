package command

import (
	"fmt"
	"os"
	"strings"

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
