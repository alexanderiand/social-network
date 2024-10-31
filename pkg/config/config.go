package config

import (
	"errors"
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

var ErrFilePathNotExists = errors.New("error, file path not exists")

type (
	Config struct {
		*Service    `yaml:"service"`
		*HTTPServer `yaml:"http_server"`
		*Cache
		*Database
	}

	Service struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTPServer struct { // host, port - это обычно хранится в переменных окружения, если это ком.разработка. стракт-тэг env-require:"true"` требует обязательно значение в env file
		Host          string        `env:"HTTP_SERVER_HOST" env-require:"true"`
		Port          string        `env:"HTTP_SERVER_PORT" env-require:"true"`
		Idletimeout   time.Duration `yaml:"idle_timeout"`
		Writretimeout time.Duration `yaml:"write_timeout"`
		Readtimeout   time.Duration `yaml:"read_timeout"`
		MaxMB         int           `yaml:"max_header_mb"`
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
