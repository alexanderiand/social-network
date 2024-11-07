package platform

import (
	"context"
	"log/slog"
	"os"

	"social-network/internal/infrastructure/repository/storage/postgresql"
	"social-network/pkg/config"
)

// TODO: Run platform method, run all sub services
// Run Platform with of the platform microservices, receive ctx, and *cfg
// If a *Config struct pointer is nil, return ErrNilStructPointer
// If the *Config params is invalid, return ErrInvalidCfgParam
// If error accrued while initializing the databases new instance, return ErrDBInit*(postgres or other db error)
// If error happened while implement DI, return more specific Err*(layer init error)
// If going unknown error, return this unknown error
func Run(ctx context.Context, cfg *config.Config) error {
	// init databases, message brokers
	// TODO: InitPostgreSQL
	postgres, err := postgresql.New(ctx, cfg)
	if err != nil {
		return err // TODO: call platform.Stop
	}

	if err := postgres.Ping(ctx); err != nil {
		return err // TODO: call platform.Stop
	}
	slog.Info("Successful connected to the PostgreSQL")

	// TODO: InitMongoDB

	// TODO: InitRedis

	// TODO: InitRabbitMQ

	//! every microservices implement into itself Dependency Inversion, use Dependency Injection

	// TODO: eventsrv.Run

	// TODO: ssosrv.Run

	// TODO: contentsrv.Run

	// TODO: chatsrv.Run

	// TODO: comsrv.Run

	return nil
}

// TODO: Stop (graceful shutdown) platform with all sub-services (micro-services)
// Stop called if context canceled, receive os.Signal, or critical error, or invalid config params
// Stop implement Graceful Shutdown for the Platform, also for every sub services of the Platform
// If ctx, sigChan, or crtErrChan is invalid, return Err*(above arg name)
// If the databases closing with error, return this specific error as a wrapped error with more info
// If the message broker, RabbitMQ return error in closing, return this error as a wrapped error with more error
// if another going wrong with unknown error, return this error as the wrapped error, with additional information
func Stop(ctx context.Context, cfg *config.Config, sigChan <-chan os.Signal, critErrChan <-chan error) error {
	// receive context canceling, os signal or critical error, start graceful shutdown the Platform

	// start closing transport layer connection (http, websocket, grpc, amqp)

	// send to the every microservices signal about starting graceful shutdown process

	// waiting closing this services

	// close the all databases and message broker
	return nil
}
