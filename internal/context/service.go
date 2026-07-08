package context

import (
	"fmt"

	"github.com/hynexis/heos-cli/internal/task"
)

type Service struct {
	AIRoot string
}

func NewService(aiRoot string) *Service {
	return &Service{
		AIRoot: aiRoot,
	}
}

func (s *Service) Generate(taskID string) error {

	// Load task
	t, err := task.Load(taskID)
	if err != nil {
		return err
	}

	// Load state
	state, err := task.LoadState(taskID)
	if err != nil {
		return err
	}

	// Resolve role
	role := ResolveRole(state)

	fmt.Println("Task Loaded :", taskID)
	fmt.Println("Current Role:", role)

	// Select context documents
	selector := NewSelector(s.AIRoot)

	files, err := selector.Select(role)
	if err != nil {
		return err
	}

	// Build context
	builder := NewBuilder()

	if err := builder.AddFiles(files); err != nil {
		return err
	}

	// Task request selalu paling akhir
	if err := builder.AddFile(t.RequestPath); err != nil {
		return err
	}

	content := builder.Build()

	if err := Write(content); err != nil {
		return err
	}

	fmt.Println("Context generated.")

	return nil
}
