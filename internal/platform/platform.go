package platform

import (
	"context"
	"log/slog"
	"os"
	"social-network/internal/infrastructure/repository/storage/postgres"
	"social-network/pkg/config"
)

// platfrom.Run
func Run(ctx context.Context, cfg *config.Config) error {
	// TODO: init databases (postgres, mongodb, redis) and message brokers (rabbitmq)

	// TODO: NewPostgres client

	slog.Debug("Start connecting to the PostgreSQL database")
	postgres, err := postgres.New(ctx, cfg)
	if err != nil {
		return err
	}
	slog.Info("Successful connection to postgres")

	if err := postgres.Ping(ctx); err != nil {
		return err
	}
	slog.Info("Successful checked connection to the PostgreSQL")

	// TODO: ssosrv.Run

	// TODO: contentsrv.Run

	// TODO: chatsrv.Run

	// TODO: comsrv.Run (com - communication, srv - service)

	// TODO: eventsrv.Run

	// TODO: if happened error, call platform.Stop
	return nil
}

// platfrom.Stop - for implementing graceful shutdown the platform
func Stop(ctx context.Context, sigChan <-chan os.Signal, crtErrChan <-chan error) error {

	return nil
}
