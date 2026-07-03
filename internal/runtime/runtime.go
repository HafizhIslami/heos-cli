package runtime

import (
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
	return r.Dispatch()
}
