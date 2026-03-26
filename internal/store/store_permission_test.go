//go:build !windows

package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStore_Save_Permission(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "store.enc")
	testPassphrase := []byte("testpassphrase")

	s := &Store{path: testFile}

	s.Secrets = map[string]string{
		"OPENAI_KEY": "my-openai-key",
		"AWS_KEY":    "my-aws-key",
	}

	err := s.Save(testPassphrase)

	if err != nil {
		t.Fatalf("expected not error got: %v", err)
	}

	info, err := os.Stat(testFile)
	if err != nil {
		t.Fatalf("failed to stat file: %v", err)
	}

	if info.Mode().Perm() != 0o600 {
		t.Errorf("want permission 0o600, got %v", info.Mode().Perm())
	}

}
