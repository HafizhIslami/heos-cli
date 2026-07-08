package runtime

import (
	"fmt"
	"os"

	"github.com/hynexis/heos-cli/internal/context"
)

func (r *Runtime) Start() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("task id required")
	}

	taskID := os.Args[2]

	service := context.NewService(".ai")

	return service.Generate(taskID)
}
