package store

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStore_Load(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		wantCount   int
	}{
		{
			name:        "empty list",
			fileContent: `{ "secrets": {} }`,
			wantCount:   0,
		},
		{
			name:        "one entry",
			fileContent: `{ "secrets": {"BD_PASSWORD": "randomlistofcharacter"} }`,
			wantCount:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			testFile := filepath.Join(tempDir, "secrets.json")

			err := os.WriteFile(testFile, []byte(tt.fileContent), 0600)
			if err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}

			s, err := Load(testFile)

			if err != nil {
				t.Fatalf("expected no error got: %v", err)
			}

			if s == nil {
				t.Fatal("expected non-nil store")
			}

			if len(s.Secrets) != tt.wantCount {
				t.Errorf("expected %d colleagues, got %d", tt.wantCount, len(s.Secrets))
			}

		})
	}
}

func TestStore_LoadError(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
	}{
		{
			name:        "empty file",
			fileContent: "",
		},
		{
			name:        "malformed json",
			fileContent: `{ "secrets": {`,
		},
		{
			name:        "invalid json",
			fileContent: `not json at all`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			testFile := filepath.Join(tempDir, "secrets.json")

			err := os.WriteFile(testFile, []byte(tt.fileContent), 0600)
			if err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}

			s, err := Load(testFile)

			if s != nil {
				t.Error("expected nil list of secrets on error")
			}

		})
	}
}
