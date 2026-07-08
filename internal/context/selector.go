package context

import (
	"fmt"
	"path/filepath"
)

type Selector struct {
	root string
}

func NewSelector(root string) *Selector {
	return &Selector{
		root: root,
	}
}

var roleDocuments = map[Role][]string{
	Architect: {
		"core/MISSION_CONTROL.md",
		"core/AGENTS.md",
		"core/ROLES.md",
		"core/WORKFLOW.md",
		"core/DECISION_ENGINE.md",

		"standards/ARCHITECTURE_STANDARD.md",

		"templates/ARCHITECTURE_PROPOSAL.md",
	},

	Engineer: {
		"core/MISSION_CONTROL.md",
		"core/AGENTS.md",
		"core/ROLES.md",
		"core/WORKFLOW.md",
		"core/ENGINEERING_PRINCIPLES.md",
		"core/EXECUTION_ORCHESTRATOR.md",

		"standards/SECURITY_STANDARD.md",

		"templates/IMPLEMENTATION_PLAN.md",
	},

	Reviewer: {
		"core/MISSION_CONTROL.md",
		"core/AGENTS.md",
		"core/ROLES.md",
		"core/WORKFLOW.md",
		"core/HANDOFF_PROTOCOL.md",

		"standards/CODE_REVIEW_STANDARD.md",
		"standards/SECURITY_STANDARD.md",

		"templates/CODE_REVIEW_REPORT.md",
	},
}

func (s *Selector) Select(role Role) ([]string, error) {

	docs, ok := roleDocuments[role]
	if !ok {
		return nil, fmt.Errorf("unknown role: %s", role)
	}

	files := make([]string, 0, len(docs))

	for _, doc := range docs {
		files = append(files, filepath.Join(s.root, doc))
	}

	return files, nil
}