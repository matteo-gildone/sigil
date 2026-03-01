package xdg

import (
	"fmt"
	"os"
	"path/filepath"
)

func DataDir() (string, error) {
	if os.Getenv("XDG_DATA_HOME") == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("DataDir(): failed to get user home directory: %w", err)
		}
		return filepath.Clean(homeDir) + "/.local/share/sigil", nil
	}

	return os.Getenv("XDG_DATA_HOME") + "/.local/share/sigil", nil
}

func ConfigDir() (string, error) {
	if os.Getenv("XDG_CONFIG_HOME") == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("ConfigDir(): failed to get user home directory: %w", err)
		}
		return filepath.Clean(homeDir) + "/.config/sigil", nil
	}

	return os.Getenv("XDG_CONFIG_HOME") + "/.config/sigil", nil
}

func ProjectPath(project string) (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return "", fmt.Errorf("ProjectPath(): failed to get data directory: %w", err)
	}
	return fmt.Sprintf("%s/%s", dataDir, project), nil
}
