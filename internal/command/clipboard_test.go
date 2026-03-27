package command

import "testing"

type mockWriter struct {
	written []byte
	err     error
}

func (m *mockWriter) Write(data []byte) error {
	m.written = data
	return m.err
}

func TestCopyToClipboard(t *testing.T) {
	w := &mockWriter{}
	err := copyToClipboard(w, "MY_KEY", "my-secret", 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(w.written) != "my-secret" {
		t.Errorf("want %q, got %q", "my-secret", string(w.written))
	}
}

func TestCopyToClipboard_Clear(t *testing.T) {
	w := &mockWriter{}
	err := copyToClipboard(w, "MY_KEY", "my-secret", 1)

	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}

	if len(w.written) != 0 {
		t.Errorf("expected clipboard to be cleared, got: %q", string(w.written))
	}
}

func TestCopyToClipboard_Error(t *testing.T) {
	err := copyToClipboard(&execWriter{
		tool: "non-existing-tool",
	}, "MY_KEY", "my-secret", 0)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
