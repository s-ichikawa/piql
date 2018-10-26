package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Env struct {
	PiqlPort       string `envconfig:"PIQL_PORT" default:"8080"`
	PixelaEndpoint string `envconfig:"PIXCELA_ENDPOINT" default:"https://pixe.la"`
}

func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}

	return &env, nil
}
