package cli

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/term"
)

var ErrNotTerminal = errors.New("passphrase input requires an interactive terminal")

func PromptPassphrase(prompt string, fd int) ([]byte, error) {
	if !term.IsTerminal(fd) {
		return nil, ErrNotTerminal
	}

	fmt.Fprint(os.Stderr, prompt)

	passphrase, err := term.ReadPassword(fd)
	if err != nil {
		return nil, fmt.Errorf("failed to read password, %w", err)
	}

	fmt.Fprintln(os.Stderr)

	return passphrase, nil
}
