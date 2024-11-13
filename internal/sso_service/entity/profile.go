package ssoentity

// Profile of the User
type Profile struct {
	ID int `json:"id" sql:"id"` 
	UserID int `json:"user_id" sql:"user_id"` 
	UserInfo string `json:"user_info,omitempty" sql:"user_info"` 
	IsPrivate bool `json:"is_private" sql:"is_private"` 
	FollowingsIDs []int `json:"followings_ids,omitempty" sql:"followings_ids"`
	FollowersIDs []int `json:"followers_ids,omitempty" sql:"followers_ids"`
	ProfilePosts []interface{} `json:"profile_posts,omitempty" sql:"profile_posts"` // profile posts IDs
}