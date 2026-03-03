package cli

import (
	"errors"
	"os"
	"testing"
)

func TestPromptPassphrase(t *testing.T) {
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	_, err := PromptPassphrase("", int(r.Fd()))

	if err == nil {
		t.Fatal("expected error when stdin is not a terminal")
	}

	if !errors.Is(err, ErrNotTerminal) {
		t.Errorf("expected %v, got %v", ErrNotTerminal, err)
	}

}
