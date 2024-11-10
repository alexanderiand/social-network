package repository

// Storage ->>> Postgres
type Storage interface {
	// TODO: Storage methods
}

type Repository struct {
	Storage // Dependency Injection
}

// New is Repository constructor, return *Repository
func New(st Storage) *Repository {
	return &Repository{Storage: st} // DI
}
