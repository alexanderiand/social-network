package logger

import (
	"errors"
	"log/slog"
	"os"
	"social-network/pkg/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var ErrNilStructPointer = errors.New("error, config structur is nil")

// InitLogger initialization logger for project
func InitLogger(cfg *config.Config) (log *slog.Logger, err error) {
	if cfg == nil {
		return nil, ErrNilStructPointer
	}

	switch cfg.Env {
	case envLocal:
		log = newCustomLogger(envLocal) // with level - debug, handler - text
	case envDev:
		log = newCustomLogger(envDev) // with level - debug, handler - json
	case envProd: // На продакшене уровень дебага не нужен, т.к это уже продакшн код
		log = newCustomLogger(envProd) // with info level, and json handler
	default:
		log = newCustomLogger(envDev) // with level info - json handler
	}

	return log, nil
}

// newCustomLogger
func newCustomLogger(level string) (logger *slog.Logger) {
	if level == "local" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		slog.SetDefault(logger)
	}

	if level == "dev" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		slog.SetDefault(logger)
	}

	if level == "prod" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		slog.SetDefault(logger)
	}

	return logger
}
