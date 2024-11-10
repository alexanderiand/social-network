package infras

import (
	"social-network/pkg/infras/cache/redis"
	"social-network/pkg/infras/message_broker/rabbitmq"
	"social-network/pkg/infras/storage/mongodb"
	"social-network/pkg/infras/storage/postgresql"
)

// Infras
type Infras struct {
	PostgreSQL *postgresql.Postgres
	MongoDB    *mongodb.MongoDB
	Redis      *redis.Redis
	RabbitMQ   *rabbitmq.RabbitMQ
}
