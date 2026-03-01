package xdg

import "testing"

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
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)
	t.Setenv("XDG_DATA_HOME", "")
	want := tempHome + "/.local/share/sigil"
	got, err := DataDir()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}
	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
