package runtime

import (
	"fmt"

	"github.com/hynexis/heos-cli/internal/config"
)

type Runtime struct {
	Config *config.Config
}

func New() (*Runtime, error) {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		return nil, err
	}

	return &Runtime{
		Config: cfg,
	}, nil
}

func (r *Runtime) Run() error {

	fmt.Println("===================================")
	fmt.Println(" HEOS CLI v1")
	fmt.Println("===================================")

	fmt.Println("Architect :", r.Config.Workers.Architect)
	fmt.Println("Engineer  :", r.Config.Workers.Engineer)

	fmt.Println("Ready.")

	return nil
}