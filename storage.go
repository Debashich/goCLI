package main

import (
	"encoding/json"
	"os"
)

type storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *storage[T] {
	return &storage[T]{FileName: fileName}
}

func (s *storage[T]) Save(data []T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *storage[T]) Load() ([]T, error) {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []T{}, nil
		}
		return nil, err
	}
	var data []T
	if err:= json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}
	return data, nil
}