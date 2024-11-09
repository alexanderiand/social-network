package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"

	"social-network/pkg/config"
	"social-network/pkg/utils"
)

const (
	_attempts = 3
	_delay    = 3 * time.Second
)

var (
	ErrNilStructPointer           = errors.New("error, nil struct pointer")
	ErrCannotConnectionToRabbitMQ = errors.New("error, can't connected to the RabbitMq")
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
	)

	err := utils.DoWithTries(func() error {
		conn, err := amqp.Dial(dsn)
		if err != nil {
			return err
		}
		rmb.Connection = conn

		return nil
	}, _attempts, _delay)

	if err != nil {
		return nil, err
	}

	return rmb, nil
}
