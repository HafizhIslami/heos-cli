package runtime

import (
	"fmt"
	"os"
)

func (r *Runtime) Dispatch() error {

	if len(os.Args) < 2 {
		return r.Help()
	}

	switch os.Args[1] {

	case "init":
		return r.Init()

	case "start":
		return r.Start()

	case "status":
		return r.Status()

	default:
		return fmt.Errorf("unknown command: %s", os.Args[1])
	}
}