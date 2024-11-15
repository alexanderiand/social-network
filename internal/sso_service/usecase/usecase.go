package ssousecase

// SSORepository is the main SSO Service Repository
type SSORepository interface {
	// TODO: repo methods
}

// Main SSO Service UseCase
type UseCase struct {
	SSOAuthUseCase
	SSORepository // Dependency Injection
}

// New is the main sso service usecase constructor
// return a new instance of the *UseCase
func New(rp SSORepository) *UseCase {
	return &UseCase{SSORepository: rp} // DI
}
