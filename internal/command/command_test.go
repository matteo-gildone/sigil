package command

import "testing"

func TestCommand_Narg(t *testing.T) {
	tests := []struct {
		name      string
		fn        func([]string) error
		wantedErr string
	}{
		{
			name:      "runSet",
			fn:        runSet,
			wantedErr: "usage: sigil set [-project] KEY",
		},
		{
			name:      "runGet",
			fn:        runGet,
			wantedErr: "usage: sigil get [-project] [-clip] [-clear 15] KEY",
		},
		{
			name:      "runDelete",
			fn:        runDelete,
			wantedErr: "usage: sigil delete [-project] KEY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn([]string{})

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != tt.wantedErr {
				t.Errorf("want: %q, got: %q", tt.wantedErr, err.Error())
			}
		})
	}
}

func TestRunGet_MissingClipboardTool(t *testing.T) {
	t.Setenv("SIGIL_CLIPBOARD", "")
	err := runGet([]string{"SOME_KEY"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err.Error() != "no clipboard tool configured: set SIGIL_CLIPBOARD or use -clip flag" {
		t.Errorf("want: %q, got: %q", "no clipboard tool configured: set SIGIL_CLIPBOARD or use -clip flag", err.Error())
	}
}
