package app

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Config struct {
	Addr string // example :9000
}

func DecodeConfig(fPath string) (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(fPath, &conf); err != nil {
		return conf, errors.WithStack(err)
	}
	return conf, nil
}
