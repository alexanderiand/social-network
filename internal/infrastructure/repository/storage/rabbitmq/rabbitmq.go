package rabbitmq

import (
	"context"
	"errors"
	"fmt"

	"github.com/streadway/amqp"

	"social-network/pkg/config"
)

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// MBTX - MessageBroker TX
type MBTX interface {
	// TODO: RabbitMQ methods
}

// RabbitMQ is a main Message Broker of the Platform
type RabbitMQ struct {
	*amqp.Connection
}

// New is constructor, return *RabbitMQ
func New(ctx context.Context, cfg *config.Config) (*RabbitMQ, error) {
	if ctx == nil || cfg == nil {
		return nil, ErrNilStructPointer
	}
	rmb := &RabbitMQ{}

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.MessageBroker.Username,
		cfg.MessageBroker.Password,
		cfg.MessageBroker.Host,
		cfg.MessageBroker.Port,
	) // TODO: fix this, change user, pass place
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	rmb.Connection = conn

	return rmb, nil
}

// TODO: doWithTries
