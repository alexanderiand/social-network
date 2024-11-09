package postgresql

import (
	"context"

	"social-network/pkg/config"
)

type DBEngine interface {
	Connection(ctx context.Context, cfg *config.Config) (*Postgres, error)
	RunMigration(path string) error
}
