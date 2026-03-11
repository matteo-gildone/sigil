package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
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

func TestStore_LoadUnexistingFile(t *testing.T) {
	s, err := Load("test-file.json")

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	if s == nil {
		t.Fatal("expected non-nil store")
	}

	if len(s.Secrets) != 0 {
		t.Errorf("expected %d secrets, got %d", 0, len(s.Secrets))
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
		"AWS_KEY": "my-aws-key",
	}}

	v, ok := s.Get("AWS_KEY")

	if v != "my-aws-key" {
		t.Errorf("want: %q, got: %q", "my-aws-key", v)
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

func TestStore_List(t *testing.T) {
	s := &Store{
		Secrets: map[string]string{
			"OPENAI_KEY": "my-openai-key",
			"AWS_KEY":    "my-aws-key",
		},
	}

	want := []string{"AWS_KEY", "OPENAI_KEY"}

	got := s.List()

	if !slices.Equal(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestStore_Delete(t *testing.T) {
	tests := []struct {
		name           string
		key            string
		expectedLength int
	}{
		{
			name:           "delete existing key",
			key:            "OPENAI_KEY",
			expectedLength: 1,
		},
		{
			name:           "delete non-existing key",
			key:            "RANDOM_KEY",
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				Secrets: map[string]string{
					"OPENAI_KEY": "my-openai-key",
					"AWS_KEY":    "my-aws-key",
				},
			}

			s.Delete(tt.key)

			if len(s.List()) != tt.expectedLength {
				t.Errorf("want: %d, got: %d", 1, len(s.List()))
			}
		})
	}
}

func TestStore_Save(t *testing.T) {
	tests := []struct {
		name           string
		secrets        map[string]string
		expectedLength int
	}{
		{
			name:           "empty key list",
			secrets:        map[string]string{},
			expectedLength: 0,
		},
		{
			name: "single key",
			secrets: map[string]string{
				"OPENAI_KEY": "my-openai-key",
			},
			expectedLength: 1,
		},
		{
			name: "multiple keys",
			secrets: map[string]string{
				"OPENAI_KEY": "my-openai-key",
				"AWS_KEY":    "my-aws-key",
			},
			expectedLength: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			testFile := filepath.Join(tempDir, "secrets.json")

			s := &Store{path: testFile}
			s.Secrets = tt.secrets

			err := s.Save()

			if err != nil {
				t.Fatalf("expected not error got: %v", err)
			}

			// Verify file was created
			if _, err := os.Stat(testFile); err != nil {
				t.Fatalf("expected file to exist: %v", err)
			}

			// Verify content
			data, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}

			var loaded Store
			err = json.Unmarshal(data, &loaded)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if len(loaded.Secrets) != tt.expectedLength {
				t.Errorf("expected %d colleagues, got %d", tt.expectedLength, len(loaded.Secrets))
			}
		})
	}

	t.Run("overwrites existing file", func(t *testing.T) {
		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "secrets.json")

		s := &Store{path: testFile, Secrets: map[string]string{}}
		s.Set("OPENAI_KEY", "my-openai-key")

		err := s.Save()

		if err != nil {
			t.Fatalf("expected not error got: %v", err)
		}

		// Verify file was created
		if _, err := os.Stat(testFile); err != nil {
			t.Fatalf("expected file to exist: %v", err)
		}

		// Verify content
		data, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		var loaded Store
		err = json.Unmarshal(data, &loaded)

		if len(loaded.Secrets) != 1 {
			t.Errorf("expected %d keys, got %d", 1, len(loaded.Secrets))
		}

		s.Set("AWS_KEY", "my-aws-key")

		err = s.Save()

		if err != nil {
			t.Fatalf("expected not error got: %v", err)
		}

		data, err = os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		err = json.Unmarshal(data, &loaded)

		if len(loaded.Secrets) != 2 {
			t.Errorf("expected %d keys, got %d", 2, len(loaded.Secrets))
		}
	})
}

func TestStore_RoundTrip(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "secrets.json")

	err := os.WriteFile(testFile, []byte("{ \"secrets\": {} }"), 0600)
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

	s.Set("AWS_KEY", "my-aws-key")
	err = s.Save()

	if err != nil {
		t.Fatalf("expected not error got: %v", err)
	}

	s, err = Load(testFile)

	if err != nil {
		t.Fatalf("expected no error got: %v", err)
	}

	v, ok := s.Get("AWS_KEY")

	if !ok {
		t.Error("expected to find the key stored")
	}

	if v != "my-aws-key" {
		t.Errorf("want: %q, got: %q", "my-aws-key", v)
	}
}
