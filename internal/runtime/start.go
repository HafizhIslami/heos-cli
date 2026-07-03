package runtime

import (
	"fmt"
	"os"

	"github.com/hynexis/heos-cli/internal/context"
	"github.com/hynexis/heos-cli/internal/task"
)

func (r *Runtime) Start() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("task id required")
	}

	taskID := os.Args[2]

	t, err := task.Load(taskID)
	if err != nil {
		return err
	}

	builder := context.NewBuilder()

	err = builder.AddFile(t.RequestPath)
	if err != nil {
		return err
	}

	content := builder.Build()

	err = context.Write(content)
	if err != nil {
		return err
	}

	fmt.Println("Context generated.")

	return nil
}