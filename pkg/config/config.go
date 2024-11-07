package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	cfgFilePathEnv = "CONFIG_FILE_PATH"

	HTTPServerHostEnv = "HTTP_SERVER_HOST"
	HTTPServerPort    = "HTTP_SERVER_PORT"

	redisURLEnv = "REDIS_URL"

	postgresHostEnv     = "POSTGRES_HOST"
	postgresPortEnv     = "POSTGRES_PORT"
	postgresUserEnv     = "POSTGRES_USER"
	postgresPasswordEnv = "POSTGRES_PASSWORD"
	postgresDBNameEnv   = "POSTGRES_DBNAME"
	postgresSSLModeEnv  = "POSTGRES_SSLMODE"

	mongoURLEnv = "MONGO_URL"
)

var (
	ErrFilePathNotExists = errors.New("error, file path not exists")
	ErrEnvNotExists      = errors.New("error, env not exists")
)

type (
	Config struct {
		Env        string `yaml:"env"`
		Service    `yaml:"service"`
		HTTPServer `yaml:"http_server"`
		Cache
		Database
	}

	Service struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTPServer struct {
		Host         string        `env:"HTTP_SERVER_HOST" env-require:"true"`
		Port         string        `env:"HTTP_SERVER_PORT" env-require:"true"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		MaxMB        int           `yaml:"max_header_mb"`
	}

	Cache struct {
		URL string `env:"REDIS_URL"` // Here we can write URL or host && port
	}

	Database struct {
		Host     string `env:"POSTGRES_HOST"`
		Port     string `env:"POSTGRES_PORT"`
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		DBname   string `env:"POSTGRES_DBNAME"`
		SSLMode  string `env:"POSTGRES_SSLMODE"`
		MongoURL string `env:"MONGO_URL"`
	}
)

// InitConfig Initialization config for project
func InitConfig() (*Config, error) {
	cfgFilePath, ok := os.LookupEnv(cfgFilePathEnv)
	fmt.Printf("Path to config: %v\n\n\n", cfgFilePath)
	if !ok {
		return nil, ErrFilePathNotExists
	}

	if _, err := os.Stat(cfgFilePath); os.IsNotExist(err) {
		return nil, err
	}

	Config := &Config{}

	if err := cleanenv.ReadConfig(cfgFilePath, Config); err != nil {
		return nil, err
	}

	return Config, nil
}
