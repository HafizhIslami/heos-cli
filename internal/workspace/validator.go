package workspace

import (
	"fmt"
	"os"
)

var required = []string{
	".ai",
	"tasks",
	"config.yaml",
}

func (w *Workspace) Validate() error {

	for _, path := range required {

		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("required path not found: %s", path)
		}

	}

	return nil
}
