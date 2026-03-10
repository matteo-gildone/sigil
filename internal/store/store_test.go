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
				t.Errorf("expected %d secrets, got %d", tt.wantCount, len(s.Secrets))
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

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if s != nil {
				t.Error("expected nil list of secrets on error")
			}

		})
	}
}

func TestStore_Set(t *testing.T) {
	tests := []struct {
		name           string
		initialStore   *Store
		newKeys        map[string]string
		expectedLength int
	}{
		{
			name:         "empty key list",
			initialStore: &Store{Secrets: map[string]string{}},
			newKeys: map[string]string{
				"new-key": "new value",
			},
			expectedLength: 1,
		},
		{
			name: "keys already present",
			initialStore: &Store{Secrets: map[string]string{
				"old-key": "old value",
			}},
			newKeys: map[string]string{
				"new-key": "new value",
			},
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for key, value := range tt.newKeys {
				tt.initialStore.Set(key, value)
			}

			if len(tt.initialStore.Secrets) != tt.expectedLength {
				t.Errorf("want: %d, got: %d", tt.expectedLength, len(tt.initialStore.Secrets))
			}

		})
	}
}

func TestStore_Get(t *testing.T) {
	s := &Store{Secrets: map[string]string{
		"AWS_KEY": "sddfdgdfgfghghgf",
	}}

	v, ok := s.Get("AWS_KEY")

	if v != "sddfdgdfgfghghgf" {
		t.Errorf("want: %q, got: %q", "sddfdgdfgfghghgf", v)
	}

	if !ok {
		t.Errorf("expected true, got: %T", ok)
	}
}

func TestStore_GetNoKey(t *testing.T) {
	s := &Store{Secrets: map[string]string{}}

	v, ok := s.Get("AWS_KEY")

	if v != "" {
		t.Errorf("expected empty string, got: %q", v)
	}

	if ok {
		t.Errorf("expected false, got: %T", ok)
	}
}
