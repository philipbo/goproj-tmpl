package app

import (
	"fmt"
	"github.com/pkg/errors"
	"path"
	"sync"
)

var (
	appOnce sync.Once
	app     *App
)

type App struct {
	Flags  Flags
	Config Config
}

func GetApp() *App {
	appOnce.Do(func() {
		app = &App{}
	})

	return app
}

func (a *App) Init() error {
	if err := a.parseFlags(); err != nil {
		return errors.WithStack(err)
	}

	fpath := path.Join(a.Flags.ConfigPath, fmt.Sprintf("%s.toml", a.Flags.Env))
	if err := a.initConfig(fpath); err != nil {
		return errors.WithStack(err)
	}

	if a.Flags.Env != "" {
		a.Config.Env = a.Flags.Env
	}

	return nil
}

func (a *App) parseFlags() error {
	if flags, err := parseFlags(); err != nil {
		return errors.Wrap(err, "parse flag err")
	} else {
		a.Flags = flags
	}
	return nil
}

func (a *App) initConfig(fpath string) error {
	if conf, err := DecodeConfig(fpath); err != nil {
		return errors.Wrapf(err, "decode config file err, fpath: %s", fpath)
	} else {
		a.Config = conf
	}

	return nil
}

//func (a *App)
