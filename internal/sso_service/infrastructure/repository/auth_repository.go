package ssorepository

import (
	ssoentity "social-network/internal/sso_service/entity"
)

type SSOAuthStorage interface {
	UserInf
	TokenInf
	SessionInf
}

// TODO: DELETE Bellow interfaces
// EventNotifierInf

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

// SSOAuthRepo
type SSOAuthRepository struct {
	db SSOAuthStorage
}

// New create a new SSOAuthRepo
func NewSSOAuthRepository(st SSOAuthStorage) *SSOAuthRepository {
	return &SSOAuthRepository{db: st}
}

// TODO: Implement bellow METHODS

func (r *SSOAuthRepository) SaveUser(user *ssoentity.User) (userID int, err error) {
	return 0, nil
}

func (r *SSOAuthRepository) GetAllUser() (users []*ssoentity.User, err error) {
	return nil, nil
}

func (r *SSOAuthRepository) GetUsersByIDRange(idRange ...int) (users []*ssoentity.User, err error) {
	return nil, nil
}

func (r *SSOAuthRepository) GetUserByID(userID int) (user *ssoentity.User, err error) {
	return nil, nil
}

func (r *SSOAuthRepository) UpdateUserByID(user *ssoentity.User) (userID int, err error) {
	return 0, nil
}

func (r *SSOAuthRepository) DeleteUserByID(userID int) error {
	return nil
}

func (r *SSOAuthRepository) SaveSession(ss *ssoentity.Session, userID int) (sessionID string, err error) {
	return "", nil
}

func UpdateSessionByUserID(userID int, ss *ssoentity.Session) error {
	return nil
}

func DeleteSessionByUserID(userID int) error {
	return nil
}

func TokenSave(token *ssoentity.JWTToken, userID int) (tokenID string, err error) {
	return "", nil
}

func GetTokenByUserID(userID int) (token *ssoentity.JWTToken, err error) {
	return nil, nil
}

func GetTokenByTokenID(tokenID string) (token *ssoentity.JWTToken, err error) {
	return nil, nil
}

func UpdateTokenByUserID(userID int, token *ssoentity.JWTToken) error {
	return nil
}

func DeleteTokenByUserID(userID int) error {
	return nil
}
