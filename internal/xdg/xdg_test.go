package xdg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDataDirEnvDir(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "/my/random/path")
	want := "/my/random/path/sigil"
	got, err := DataDir()

	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestDataDirNoEnvDir(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "")
	homeDir, _ := os.UserHomeDir()
	want := filepath.Join(homeDir, ".local", "share", "sigil")
	got, err := DataDir()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}
	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestConfigDirEnvDir(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "/my/random/path")
	want := "/my/random/path/sigil"
	got, err := ConfigDir()

	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestConfigDirNoEnvDir(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "")
	homeDir, _ := os.UserHomeDir()
	want := filepath.Join(homeDir, ".config", "sigil")
	got, err := ConfigDir()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}
	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestProjectPath(t *testing.T) {
	dataHome := t.TempDir()
	t.Setenv("XDG_DATA_HOME", dataHome)
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			"create default",
			"default",
			filepath.Join(dataHome, "sigil", "default"),
		},
		{
			"create project 1",
			"project1",
			filepath.Join(dataHome, "sigil", "project1"),
		},
		{
			"create project 2",
			"project2",
			filepath.Join(dataHome, "sigil", "project2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProjectPath(tt.input)
			if err != nil {
				t.Fatalf("expected no error got: %v", err)
			}

			if _, err := os.Stat(got); err != nil {
				t.Fatalf("expected path to exist: %v", err)
			}

			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}
