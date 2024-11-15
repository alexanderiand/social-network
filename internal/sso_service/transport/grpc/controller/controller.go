package grpccontroller

// SSOUsecase
type SSOUseCase interface {
	SSOAuthUseCase
}

// SSOAuthInfc

// SSOController gRPC controller of the SSO Service
type SSOController struct {
	Auth *SSOAuthController
}

// New create a new SSOController instance
func New(uc SSOUseCase) *SSOController {
	return &SSOController{
		Auth: NewSSOAuthController(uc),
	}
}
