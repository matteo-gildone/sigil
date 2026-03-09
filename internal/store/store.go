package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
)

type Store struct {
	Secrets map[string]string `json:"secrets"`
	path    string
}

func Load(path string) (*Store, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Store{path: path, Secrets: map[string]string{}}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	s := &Store{path: path}

	if err := json.Unmarshal(file, s); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return s, nil
}

func (s *Store) Save() error {
	js, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, js, 0600)
}

func (s *Store) Set(key, value string) error {
	s.Secrets[key] = value
	return nil
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

	for commentType := range s.Secrets {
		keys = append(keys, commentType)
	}

	sort.Strings(keys)

	return keys
}
