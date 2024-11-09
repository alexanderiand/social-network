package platform

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"social-network/internal/infrastructure/repository/storage/mongodb"
	"social-network/internal/infrastructure/repository/storage/postgresql"
	"social-network/internal/infrastructure/repository/storage/rabbitmq"
	"social-network/internal/infrastructure/repository/storage/redis"
	"social-network/pkg/config"
)

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// Run Platform with of the platform microservices, receive ctx, and *cfg
// If a *Config struct pointer is nil, return ErrNilStructPointer
// If the *Config params is invalid, return ErrInvalidCfgParam
// If error accrued while initializing the databases new instance, return ErrDBInit*(postgres or other db error)
// If error happened while implement DI, return more specific Err*(layer init error)
// If going unknown error, return this unknown error
func Run(ctx context.Context, cfg *config.Config) error {
	// init databases, cache, and message brokers
	infra, err := InitInfrastructures(ctx, cfg)
	if err != nil {
		return err
	}
	slog.Info("Infrastructures successful initialized")

	_ = infra

	//! every microservices implement into itself Dependency Inversion, use Dependency Injection

	// srv - service
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

// Infras
type Infras struct {
	PostgreSQL *postgresql.Postgres
	MongoDB    *mongodb.MongoDB
	Redis      *redis.Redis
	RabbitMQ   *rabbitmq.RabbitMQ
}

// InitInfrastructures create the new clients and the new connections
// with databases and message brokers, return *Infras, and nil as error, if ok
// If ctx or cfg is nil, return ErrNilStructPointer
func InitInfrastructures(ctx context.Context, cfg *config.Config) (*Infras, error) {
	if ctx == nil || cfg == nil {
		return nil, ErrNilStructPointer
	}
	infra := &Infras{}

	// InitPostgreSQL
	postgres, err := postgresql.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Debug("Created a new PostgreSQL client")

	if err := postgres.Ping(ctx); err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Info("Successful connected to the PostgreSQL")

	if err := postgres.RunMigration(cfg.MigrationFilesPath); err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Info("Successful completed migrations via goose")
	defer postgres.Close()
	infra.PostgreSQL = postgres

	// InitMongoDB
	mongo, err := mongodb.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Debug("Created a new MongoDB client")

	if err := mongo.Ping(ctx, nil); err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Info("Successful connected to the MongoDB")
	defer mongo.Disconnect(ctx)
	infra.MongoDB = mongo

	// InitRedis
	redis, err := redis.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Debug("Created a new client instance of the Redis")

	if res := redis.Ping(); res.Err() != nil {
		return nil, res.Err() // TODO: call platform.Stop
	}
	slog.Info("Successful connected to the Redis cache server")
	defer redis.Close()
	infra.Redis = redis

	// InitRabbitMQ
	rabbit, err := rabbitmq.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err // TODO: call platform.Stop
	}
	slog.Debug("Created a new connection with the RabbitMQ message broker")
	defer rabbit.Close()
	infra.RabbitMQ = rabbit

	// mbErrChan := make(chan *amqp.Error)
	// notifyMBErr := rabbit.NotifyClose(mbErrChan)

	// for err := range notifyMBErr {
	// 	if err != nil {
	// 		slog.Error(err.Error())
	// 		return nil, err // TODO: call platform.Stop
	// 	}
	// }

	slog.Info("Successful connected to the RabbitMQ Message Broker")

	return infra, nil
}
