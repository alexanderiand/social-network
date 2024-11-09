package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"social-network/pkg/config"
	"social-network/pkg/utils"
)

// const
const (
	_attempts = 3
	_delay    = time.Second * 3
)

// common mongodb errors
var (
	ErrNilStructPointer     = errors.New("error, nil struct pointer")
	ErrCannotConnectToMongo = errors.New("error, cannot connect to the MongoDB")
)

// MongoDB is a MongoDB client
type MongoDB struct {
	*mongo.Client
}

// New is the MongoDB client constructor
// return a new instance of the *MongoDB
func New(ctx context.Context, cfg *config.Config) (*MongoDB, error) {
	if cfg == nil || ctx == nil {
		return nil, ErrNilStructPointer
	}

	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.Mongo.Username,
		cfg.Mongo.Password,
		cfg.Mongo.Host,
		cfg.Mongo.Port,
	)

	clientOps := options.Client().ApplyURI(dsn)
	db := &MongoDB{}

	err := utils.DoWithTries(func() error {
		mongodb, err := mongo.Connect(ctx, clientOps)
		if err != nil {
			return err
		}
		db.Client = mongodb

		return nil
	}, _attempts, _delay)

	if err != nil {
		return nil, err
	}

	return db, nil
}
