package context

import "github.com/hynexis/heos-cli/internal/task"

type Role string

const (
	Architect Role = "architect"
	Engineer  Role = "engineer"
	Reviewer  Role = "reviewer"
)

func ResolveRole(state *task.State) Role {

	switch state.Worker {

	case "architect":
		return Architect

	case "reviewer":
		return Reviewer

	default:
		return Engineer
	}
}