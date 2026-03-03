package cli

import (
	"errors"
	"os"
	"testing"
)

func TestPromptPassphrase(t *testing.T) {
	r, _, _ := os.Pipe()
	defer r.Close()

	_, err := PromptPassphrase("", int(r.Fd()))

	if err == nil {
		t.Fatal("expected error when stdin is not a terminal")
	}

	if !errors.Is(err, ErrIsNotTerminal) {
		t.Errorf("expected %v, got %v", ErrIsNotTerminal, err)
	}

}
