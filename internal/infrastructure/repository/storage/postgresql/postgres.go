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
)

// consts
const (
	_attempts = 3
	_delay    = time.Second * 3
)

// Postgres errors
var (
	ErrNilStructPointer   = errors.New("error, nil struct pointer")
	ErrCannotConnectToPgs = errors.New("error, 0 connection attempts left: the database is't connected")
)

// DBTX
type DBTX interface {
	// TODO: Database Methods
}

// TODO: add migrations

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
	err := doWithTries(func() error {
		db, err := pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		postgres.Pool = db

		return nil
	}, _attempts, _delay)

	if err != nil {
		return nil, err
	}

	return &postgres, nil
}

// TODO: RunMigration

// doWithTries
func doWithTries(fn func() error, attempts int, delay time.Duration) error {
	for attempts > 0 {
		if err := fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return ErrCannotConnectToPgs
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
