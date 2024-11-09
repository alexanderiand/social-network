package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis"

	"social-network/pkg/config"
)

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// CacheTX
type CacheTX interface {
	// TODO: Cache methods
}

// Redis is client the Redis,
// in platform represent central cache service
type Redis struct {
	*redis.Client
}

// New is constructor of the Redis Cache Server client
// return a new *Redis instance
// If ctx, or cfg params is invalid, return ErrNilStructPointer
func New(ctx context.Context, cfg *config.Config) (*Redis, error) {
	if cfg == nil || ctx == nil {
		return nil, ErrNilStructPointer
	}

	dsn := fmt.Sprintf("%s:%s",
		cfg.Cache.Host,
		cfg.Cache.Port,
	)
	clientOps := &redis.Options{Addr: dsn}
	rdb := &Redis{}
	rdb.Client = redis.NewClient(clientOps)

	return rdb, nil
}
