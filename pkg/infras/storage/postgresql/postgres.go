package postgresql

import (
	"context"
	"errors"
	"fmt"
	"time"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"

	"social-network/pkg/config"
	"social-network/pkg/utils"
)

// const
const (
	_attempts = 3
	_delay    = time.Second * 3
)

// Postgres errors
var (
	ErrNilStructPointer   = errors.New("error, nil struct pointer")
	ErrCannotConnectToPgs = errors.New("error, 0 connection attempts left: the database is't connected")
)

// Postgres is a abs to the PostgreSQL DB
type Postgres struct {
	DBEngine // postgresql/interface.go
	*pgxpool.Pool
}

// New is Postgres constructor, return a new
// instance of the Postgres
func New() *Postgres {
	return &Postgres{}
}

// Connection to the PostgreSQL, use pgxpool, return *pgxpool.Pool inside *Postgres
// If cfg is nil, return ErrNilStructPointer
func (p *Postgres) Connection(ctx context.Context, cfg *config.Config) (*Postgres, error) {
	if cfg == nil {
		return nil, ErrNilStructPointer
	}

	// DSN
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBname,
		cfg.Database.SSLMode,
	)

	err := utils.DoWithTries(func() error {
		db, err := pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		p.Pool = db

		return nil
	}, _attempts, _delay)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// RunMigration use the goose libs, set postgres dialect, and run migrations
// If when set dialect happened error, return this error
// If when running migration occurred error, return this error
func (p *Postgres) RunMigration(path string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(p.Pool)
	if err := goose.Up(db, path); err != nil {
		return err
	}

	return nil
}
