package command

import (
	"fmt"
	"os"
	"strings"

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

func getPassword() ([]byte, error) {
	password, err := cli.PromptPassphrase("password:", int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to read password: %w", err)
	}
	return password, nil
}
