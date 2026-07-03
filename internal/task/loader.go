package task

import (
	"fmt"
	"os"
	"path/filepath"
)

func Load(id string) (*Task, error) {

	dir := filepath.Join("tasks", id)

	if _, err := os.Stat(dir); err != nil {
		return nil, fmt.Errorf("task %s not found", id)
	}

	request := filepath.Join(dir, "request.md")

	if _, err := os.Stat(request); err != nil {
		return nil, fmt.Errorf("request.md not found")
	}

	return &Task{
		ID:          id,
		RequestPath: request,
		StatePath:   filepath.Join(dir, "state.json"),
	}, nil
}