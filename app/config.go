package app

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr string // example :9000
	Env  string
}

func DecodeConfig(fPath string) (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(fPath, &conf); err != nil {
		return conf, errors.WithStack(err)
	}

	log.Infof("config: %#v", conf)
	return conf, nil
}
