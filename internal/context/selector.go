package context

type Role string

const (
	RoleArchitect Role = "architect"
	RoleEngineer  Role = "engineer"
	RoleReviewer  Role = "reviewer"
)