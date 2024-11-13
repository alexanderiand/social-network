package ssousecase

import (
	"context"

	ssoentity "social-network/internal/sso_service/entity"
)

type SSOAuthURepo interface {
	SignUp(ctx context.Context, user *ssoentity.User) (userID int, err error)
	SignIn(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error
	SignOut(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error 
	RefreshToken(ctx context.Context, ss *ssoentity.Session, token *ssoentity.JWTToken) error 
}

// SSOAuthUseCase
type SSOAuthUseCase struct {
	repo SSOAuthURepo
}

// New create a new SSOAuthUsecase, return *SSOAuthUseCase
func NewSSOAuthUseCase(rp SSOAuthURepo) *SSOAuthUseCase {
	return &SSOAuthUseCase{
		repo: rp,
	}
}
