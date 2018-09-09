package app

import (
	"errors"
	"flag"
	"os"
)

type EnvMode string

var (
	DevMode  = "dev"
	TestMode = "test"
	ProdMode = "prod"
)

var (
	ErrConfigPathEmpty = errors.New("config path must be not empty")
	ErrEnvInvalid      = errors.New("application Env invalid ")
)

type Flags struct {
	ConfigPath string
	Env        string
}

func parseFlags() (Flags, error) {
	flags := Flags{}
	var path string
	flag.StringVar(&path, "config", "./config", "config file dir")
	if path == "" {
		return flags, ErrConfigPathEmpty
	}

	if _, err := os.Stat(path); err != nil {
		return flags, err
	}

	var env string
	flag.StringVar(&env, "Env", DevMode, "application Env (dev|test|prod)")
	switch env {
	case DevMode, TestMode, ProdMode:
	default:
		return flags, ErrEnvInvalid

	}

	flags.ConfigPath = path
	flags.Env = env

	return flags, nil
}
