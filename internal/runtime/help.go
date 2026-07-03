package runtime

import "fmt"

func (r *Runtime) Help() error {

	fmt.Println()

	fmt.Println("HEOS CLI")

	fmt.Println()

	fmt.Println("Usage")

	fmt.Println("  heos init")
	fmt.Println("  heos start <task>")
	fmt.Println("  heos status")

	fmt.Println()

	return nil
}