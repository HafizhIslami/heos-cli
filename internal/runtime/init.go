package runtime

import (
	"fmt"

	"github.com/hynexis/heos-cli/internal/workspace"
)

func (r *Runtime) Init() error {

	ws := workspace.New()

	fmt.Println("Validating workspace...")

	if err := ws.Validate(); err != nil {
		return err
	}

	if err := ws.PrepareContext(); err != nil {
		return err
	}

	fmt.Println("Workspace Ready")

	return nil
}