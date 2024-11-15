package ssoentity

import "time"

// Sessions
type Session struct {
	ID        int           `json:"id" sql:"id"`
	UserID    int           `json:"user_id" sql:"user_id"`
	ExpiresAt time.Duration `json:"expires_at" sql:"expires_at"`
	CreatedAt time.Time     `json:"created_at,omitempty" sql:"created_at"`
}

// JWTTokens
type JWTToken struct {
	ID           string        `json:"id" sql:"id"`
	UserID       int           `json:"user_id" sql:"user_id"`
	RefreshToken string        `json:"refresh_token" sql:"refresh_token"`
	ExpiresAt    time.Duration `json:"expires_at" sql:"expires_at"`
	CreatedAt    time.Time     `json:"created_at,omitempty" sql:"created_at"`
}
