package ssousecase

import (
	"context"
	"log/slog"

	ssoentity "social-network/internal/sso_service/entity"
)

type SSOAuthURepo interface {
	UserInf
	TokenInf
	SessionInf
}

// EventNotifierInf
type EventNotifierInf interface {
	NotifyNewUser(user *ssoentity.User, eventType, payload, status string) (eventId int, err error)
	// TODO: add more event notify methods
}

// TokenInf
type TokenInf interface {
	TokenSave(token *ssoentity.JWTToken, userID int) (tokenID string, err error)
	GetTokenByUserID(userID int) (token *ssoentity.JWTToken, err error)
	GetTokenByTokenID(tokenID string) (token *ssoentity.JWTToken, err error)
	UpdateTokenByUserID(userID int, token *ssoentity.JWTToken) error
	DeleteTokenByUserID(userID int) error
}

// SessionInf
type SessionInf interface {
	SaveSession(ss *ssoentity.Session, userID int) (sessionID string, err error)
	UpdateSessionByUserID(userID int, ss *ssoentity.Session) error
	DeleteSessionByUserID(userID int) error
}

// UserInf
type UserInf interface {
	SaveUser(user *ssoentity.User) (userID int, err error)
	GetAllUser() (users []*ssoentity.User, err error)
	GetUsersByIDRange(idRange ...int) (users []*ssoentity.User, err error)
	GetUserByID(userID int) (user *ssoentity.User, err error)
	UpdateUserByID(user *ssoentity.User) (userID int, err error)
	DeleteUserByID(userID int) error
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

// SignUp description
func (a *SSOAuthUseCase) SignUp(ctx context.Context, user *ssoentity.User) (userID int, err error) {
	slog.Error("usecase.SignUp: Implement me")

	return 0, nil
}

// SignIn
func (a *SSOAuthUseCase) SignIn(
	ctx context.Context,
	user *ssoentity.User,
	ss *ssoentity.Session,
	token *ssoentity.JWTToken,
) error {

	slog.Error("ssousecase.SignIn: implement me")
	return nil
}

// SignOut
func (a *SSOAuthUseCase) SignOut(
	ctx context.Context,
	userID int,
	ss *ssoentity.Session,
	token *ssoentity.JWTToken) error {

	slog.Error("ssousecase.SignOut: implement me")

	return nil
}

// RefreshToken
func (a *SSOAuthUseCase) RefreshToken(
	ctx context.Context,
	ss *ssoentity.Session,
	token *ssoentity.JWTToken) error {

	slog.Error("ssousecase.RefreshToken: Implement me")

	return nil
}
