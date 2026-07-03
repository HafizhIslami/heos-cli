package context

import (
	"os"
	"strings"
)

type Builder struct {
	lines []string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) AddFile(path string) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	b.lines = append(b.lines, string(data))

	return nil
}

func (b *Builder) Build() string {

	return strings.Join(b.lines, "\n\n")
}