package postgres

import (
	ssoentity "social-network/internal/sso_service/entity"
	"social-network/pkg/infras/storage/postgresql"
)

// DBTX
type DBTX interface {
	UserInf
	TokenInf
	SessionInf
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

type Postgres struct {
	*postgresql.Postgres
}

// New receive a new *Postgres, return
func New(pg *postgresql.Postgres) *Postgres {
	return &Postgres{
		Postgres: pg,
	}
}

// TODO: Implement bellow METHODS

func (p *Postgres) SaveUser(user *ssoentity.User) (userID int, err error) {
	return 0, nil
}

func (p *Postgres) GetAllUser() (users []*ssoentity.User, err error) {
	return nil, nil
}

func (p *Postgres) GetUsersByIDRange(idRange ...int) (users []*ssoentity.User, err error) {
	return nil, nil
}

func (p *Postgres) GetUserByID(userID int) (user *ssoentity.User, err error) {
	return nil, nil
}

func (p *Postgres) UpdateUserByID(user *ssoentity.User) (userID int, err error) {
	return 0, nil
}

func (p *Postgres) DeleteUserByID(userID int) error {
	return nil
}

func (p *Postgres) SaveSession(ss *ssoentity.Session, userID int) (sessionID string, err error) {
	return "", nil
}

func (p *Postgres) UpdateSessionByUserID(userID int, ss *ssoentity.Session) error {
	return nil
}

func (p *Postgres) DeleteSessionByUserID(userID int) error {
	return nil
}

func (p *Postgres) TokenSave(token *ssoentity.JWTToken, userID int) (tokenID string, err error) {
	return "", nil
}

func (p *Postgres) GetTokenByUserID(userID int) (token *ssoentity.JWTToken, err error) {
	return nil, nil
}

func (p *Postgres) GetTokenByTokenID(tokenID string) (token *ssoentity.JWTToken, err error) {
	return nil, nil
}

func (p *Postgres) UpdateTokenByUserID(userID int, token *ssoentity.JWTToken) error {
	return nil
}

func (p *Postgres) DeleteTokenByUserID(userID int) error {
	return nil
}
