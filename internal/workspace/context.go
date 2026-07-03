package workspace

import (
	"os"
	"path/filepath"
)

const contextDir = ".context"
const contextFile = "current.md"

func (w *Workspace) PrepareContext() error {

	if err := os.MkdirAll(contextDir, 0755); err != nil {
		return err
	}

	path := filepath.Join(contextDir, contextFile)

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
