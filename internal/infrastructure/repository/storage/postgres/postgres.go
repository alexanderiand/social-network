package postgres

import (
	"context"
	"errors"
	"fmt"
	"social-network/pkg/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_attempts = 3 // Такое именование обозначает что это это переменная константа и её не нужно трогать
	_delay    = 3 * time.Second
)

var (
	ErrNilStructPointer           = errors.New("error, struct nil")
	ErrCannotConnectionToPostgres = errors.New("error, can't connection to postrgres")
)

// 1. Define interface for a storage methods
type DBTX interface {
	// Here add method postgres
}

// 2. Declare struct as Postgres, inject a above interface into this struct
type Postgres struct {
	DBTX
	*pgxpool.Pool
}

// 3. Define New constructore, constructore return *Postgres struct pointer, and error
// New(ctx context.Context, cfg *config.Config) (*Postgres, error)
// a. check the ctx and cfg args - red case for this alg

// b. prepare Data Source Name - dsn (database connection uri)

// c. create a new connection with postgres via pgxpool lib
// d. error handling

// d. return the new instance of the Postgres

// postgres.New()

func New(ctx context.Context, cfg *config.Config) (*Postgres, error) {
	src := "internal.infra.reposiroty.storage.postgres.New" // Раньше в пакете slog не было проявления пути, и раньше разрабы указывали вручную источник откуда прилетела ошибка

	if ctx == nil || cfg == nil {
		return nil, fmt.Errorf("%s %w", src, ErrNilStructPointer)
	}

	// URI
	// postgresql://username:password@host:port/dbname?sslmode=sslmodeEnv %s // Это URI для подключения к БД, считай интернет адрес для подключения к БД (Data Source Name)
	dsn := fmt.Sprintf("postgers://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBname,
		cfg.Database.SSLMode,
	)

	fmt.Printf("\n\nPostgres DSN: %+v \n\n", dsn)

	postgres := &Postgres{}

	err := doWithTries(func() error {
		pool, err := pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}
		postgres.Pool = pool

		return nil
	}, _attempts, _delay)

	if err != nil {
		return nil, err
	}

	return postgres, nil
}

// 4. Define doWithTries func, for implement Retry pattern
func doWithTries(fn func() error, attempts int, delay time.Duration) error {
	for attempts > 0 {
		if err := fn(); err != nil {
			// slog.Warn
			attempts--
			time.Sleep(delay)
			continue
		}
		return nil
	}
	return ErrCannotConnectionToPostgres
}

// DBTX
// type DBTX interface {
// 	// TODO: postgres methods
// 	CreateUser(user entity.User) (id int, err error)
// }

// // Postgres
// type Postgres struct {
// 	DBTX
// }

// func (p *Postgres) CreateUser(user entiry.User) (id int, err error) {
// 	return nil, nil
// }

// /*
// package main

// type Storage interface {
// 	CreateUse(....)...
// }

// type UserCreator struct {
// 	Storage
// }

// func New(st Storage) *UserCreator {

// }

// func (u *UserCreator) CreateUser(user entity.User) (id int, err error) {}

// */
