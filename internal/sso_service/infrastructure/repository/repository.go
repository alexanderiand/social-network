package ssorepository

// Storage ->>> Postgres
type Storage interface {
	UserInf
	TokenInf
	SessionInf
}

type Repository struct {
	ARepo *SSOAuthRepository
}

// New is Repository constructor, return *Repository
func New(st Storage) *Repository {
	return &Repository{
		ARepo: NewSSOAuthRepository(st),
	} // DI
}
