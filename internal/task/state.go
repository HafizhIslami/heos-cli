package task

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type State struct {
	State    string `json:"state"`
	Worker   string `json:"worker"`
	Progress int    `json:"progress"`
}

func LoadState(taskID string) (*State, error) {

	path := filepath.Join("tasks", taskID, "state.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s State

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
