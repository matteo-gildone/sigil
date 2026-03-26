package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/matteo-gildone/sigil/internal/crypto"
)

type Store struct {
	Secrets map[string]string `json:"secrets"`
	path    string
}

func Load(path string, passphrase []byte) (*Store, error) {
	filePath := filepath.Join(path, "store.enc")
	file, err := os.ReadFile(filePath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Store{path: filePath, Secrets: map[string]string{}}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	data, err := crypto.Decrypt([]byte(passphrase), file)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt file: %w", err)
	}

	s := &Store{
		path: filePath,
	}

	if err := json.Unmarshal(data, s); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return s, nil
}

func (s *Store) Save(passphrase []byte) error {
	data, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to marshal store: %w", err)
	}

	encrypted, err := crypto.Encrypt([]byte(passphrase), data)
	if err != nil {
		return fmt.Errorf("failed to encrypt data: %w", err)
	}

	tmp, err := os.CreateTemp(filepath.Dir(s.path), ".sigil-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}

	removeTemp := true
	defer func() {
		if removeTemp {
			os.Remove(tmp.Name())
		}
	}()

	if _, err := tmp.Write(encrypted); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if err := tmp.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	if err := tmp.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	removeTemp = false
	return os.Rename(tmp.Name(), s.path)
}

func (s *Store) Set(key, value string) {
	s.Secrets[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	v, ok := s.Secrets[key]
	return v, ok
}

func (s *Store) Delete(key string) {
	delete(s.Secrets, key)
}

func (s *Store) List() []string {
	keys := make([]string, 0, len(s.Secrets))

	for key := range s.Secrets {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
