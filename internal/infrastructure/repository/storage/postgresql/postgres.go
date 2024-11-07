package postgresql

import (
	"context"
	"errors"
	"fmt"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	"social-network/pkg/config"
)

// Postgres errors
var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// DBTX
type DBTX interface {
	// TODO: Database Methods
}

// Postgres is a abs to the PostgreSQL DB
type Postgres struct {
	DBTX
	*pgxpool.Pool
}

// New is Postgres constructor, return a new
// instance of the Postgres
func New(ctx context.Context, cfg *config.Config) (*Postgres, error) {
	if cfg == nil {
		return nil, ErrNilStructPointer
	}

	// DoWithTries

	// DSN
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBname,
		cfg.Database.SSLMode,
	)

	postgres := Postgres{}
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	postgres.Pool = db

	return &postgres, nil
}
