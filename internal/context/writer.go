package context

import (
	"os"
	"path/filepath"
)

func Write(content string) error {

	err := os.MkdirAll(".context", 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Join(".context", "current.md"),
		[]byte(content),
		0644,
	)
}