package command

import (
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/cli"
	"github.com/matteo-gildone/sigil/internal/store"
	"github.com/matteo-gildone/sigil/internal/xdg"
)

func loadStore(project string) (*store.Store, []byte, error) {
	passphrase, err := cli.PromptPassphrase("passphrase:", int(os.Stdin.Fd()))
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
