package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	_localCfgFilePath = "/internal/sso_service/app/config/config.yaml"
	_localEnvFile     = ".env"
)

var (
	ErrConfig                = errors.New("error, local config error")
	ErrLocalCfgFileNotExists = errors.New("error, a local cfg file is not exists")
)

type (
	Config struct {
		Service `yaml:"service"`
		HTTP
	}
	Service struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Host string `env:"SSOSRV_HTTP_HOST" env-required:"true"`
		Port string `env:"SSOSRV_HTTP_PORT" env-required:"true"`
	}
)

// InitSSOSRVConfig read cfg params from local config.yaml, and .env
// return a new *Config instance with cfg params, and nil as error
// If local cfg file not exists, return ErrLocalCfgFileNotExists
// If can't read the cfg params cleanenv.ReadConfig, return wrapped err (ErrConfig + err)
func InitSSOSRVConfig() (*Config, error) {
	// config file path
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	cfgFilePath := wd + _localCfgFilePath

	// check existing the cfg file
	if _, err := os.Stat(cfgFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w %w", ErrLocalCfgFileNotExists, err)
	}
	cfg := &Config{}

	// read config params
	if err := cleanenv.ReadConfig(cfgFilePath, cfg); err != nil {
		return nil, fmt.Errorf("%s %w", ErrConfig, err)
	}

	return cfg, nil
}
