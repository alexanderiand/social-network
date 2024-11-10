package postgres

import "social-network/pkg/infras/storage/postgresql"

// DBTX
type DBTX interface {
	// TODO: postgres methods
}

type Postgres struct {
	*postgresql.Postgres
}

// New receive a new *Postgres, return
func New(pg *postgresql.Postgres) *Postgres {
	return &Postgres{
		Postgres: pg,
	}
}
