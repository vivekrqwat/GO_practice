package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	filename string
}

func newStrogare[T any](fname string) *Storage[T] {
	return &Storage[T]{filename: fname}
}
func (s *Storage[T]) Save(data T) error {
	value, er := json.MarshalIndent(data, "", "  ")
	if er != nil {
		return er
	}
	return os.WriteFile(s.filename, value, 0644)

}

func (s *Storage[T]) Load(data *T) error {
	value, er := os.ReadFile(s.filename)
	if er != nil {
		return er
	}
	return json.Unmarshal(value, data)
}
