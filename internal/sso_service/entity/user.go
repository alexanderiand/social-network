package ssoentity

import "time"

// User
type User struct {
	ID int `json:"id" sql:"id"`
	FirstName string `json:"first_name" sql:"first_name"` 
	LastName string `json:"last_name" sql:"last_name"` 
	DateOfBirth time.Time `json:"date_of_birth" sql:"date_of_birth"`
	Email string `json:"email" sql:"email"` 
	Avatar string `json:"avatar" sql:"avatar"`
	NickName string `json:"nickname,omitempty" sql:"nickname"` 
	AboutMe string `json:"about_me" sql:"about_me"` 
	Password string `json:"password" sql:"password_hash"` 
	GroupsIDs []int `json:"group_ids,omitempty" sql:"group_ids"`
	GroupChatsIDs []int `json:"groupchats_ids,omitempty" sql:"groupchats_ids"`
	PvChatsIDs []int `json:"pvchats_ids,omitempty" sql:"pvchats_ids"` 
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
}