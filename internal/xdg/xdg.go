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
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		return filepath.Join(homeDir, ".local", "share", "sigil"), nil
	}

	return filepath.Join(os.Getenv("XDG_DATA_HOME"), "sigil"), nil
}

func ConfigDir() (string, error) {
	if os.Getenv("XDG_CONFIG_HOME") == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		return filepath.Join(homeDir, ".config", "sigil"), nil
	}

	return filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "sigil"), nil
}

func ProjectPath(project string) (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return "", fmt.Errorf("failed to get data directory: %w", err)
	}
	path := fmt.Sprintf("%s/%s", dataDir, project)
	if err := os.MkdirAll(path, 0o700); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	return path, nil
}
