package infras

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"social-network/pkg/config"
	"social-network/pkg/infras/cache/redis"
	"social-network/pkg/infras/message_broker/rabbitmq"
	"social-network/pkg/infras/storage/mongodb"
	"social-network/pkg/infras/storage/postgresql"
)

// waiting n second before starting closing connections
const _waiting = time.Second * 3

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

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
	postgres, err := postgresql.New().Connection(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Debug("Created a new PostgreSQL client")

	// Check the PostgreSQL connection
	if err := postgres.Ping(ctx); err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Info("Successful connected to the PostgreSQL")

	// Run Migration via goose cli utils & lib
	if err := postgres.RunMigration(cfg.MigrationFilesPath); err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Info("Successful completed migrations via goose")
	infra.PostgreSQL = postgres

	// InitMongoDB
	mongo, err := mongodb.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Debug("Created a new MongoDB client")

	if err := mongo.Ping(ctx, nil); err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Info("Successful connected to the MongoDB")
	infra.MongoDB = mongo

	// InitRedis
	redis, err := redis.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Debug("Created a new client instance of the Redis")

	if res := redis.Ping(); res.Err() != nil {
		return nil, res.Err()
	}
	slog.Info("Successful connected to the Redis cache server")
	defer redis.Close()
	infra.Redis = redis

	// InitRabbitMQ
	rabbit, err := rabbitmq.New(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
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

// CloseConnections
func CloseConnections(ctx context.Context, infra *Infras) error {
	fn := "platform.CloseConnections"
	if infra == nil {
		return fmt.Errorf("%s %w", fn, ErrNilStructPointer)
	}
	time.Sleep(_waiting) // waiting n second before start closing conns

	// close the PostgreSQL connection
	infra.PostgreSQL.Close()
	slog.Debug("Close the PostgreSQL database connection")

	// close the MongoDB connection
	if err := infra.MongoDB.Disconnect(ctx); err != nil {
		return err
	}
	slog.Debug("Close the MongoDB connection")

	// close the Redis connection
	if err := infra.Redis.Close(); err != nil {
		return err
	}
	slog.Debug("Close the Redis connection")

	// Close the RabbitMQ connection
	if err := infra.RabbitMQ.Close(); err != nil {
		return err
	}
	slog.Debug("Close the RabbitMQ connection")

	return nil
}
