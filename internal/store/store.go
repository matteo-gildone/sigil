package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

//func (s *Store) Save() error                   {}
//func (s *Store) Set(key, value string) error   {}
//func (s *Store) Get(key string) (string, bool) {}
//func (s *Store) Delete(key string) error       {}
//func (s *Store) List() []string                {}
