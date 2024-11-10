package router

import (
	"context"
	"net/http"

	"social-network/internal/sso_service/transport/http/rest/controller"
)

// Handler
type Handler interface {
	Init(ctx context.Context, crl *controller.Controller)
}

// Router
type Router struct {
	Mux *http.ServeMux
	Ctl *controller.Controller
}

// New is constructor of the Router, return a new *Router instance
func New(ctl *controller.Controller) *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

// InitRouter mapped every router to the controller handler
func (r *Router) Init(ctx context.Context) {
	// TODO: create a new base middleware and run this mdl methods

	// SSO Service routes
	r.Mux.HandleFunc("GET /", r.Ctl.BaseController) // for checking

	// Auth endpoints
	// sign-up - registration // sessions and cookies based auth
	// sign-in - login
	// sign-out - logout

	// for distributed system, Bearer JWT Auth based auth
	// access-token
	// refresh-token

	// User endpoints
	// create a user
	// get all users
	// get the user by id
	// update the user by id
	// delete the user by id

	// Roles endpoints
	// make user group moderator by user_id for group_id
	// revoke moderator role by user_id for group_id
	// make user group chat moderator by user_id for group_chat_id
	// revoke moderator role by user_id for group_chat_id

	// Profile endpoints
	// create a user profile
	// get the user profile by user_id
	// create post for user profile by profile_id
	// make_user_profile_private by user_id, profile_id
	// make_user_profile_public by user_id, profile_id
	// get all the user followings by user_id, profile_id
	// get all the user followers by user_di, profile_id

	// log information
	// Print Routes via fmt.Printf for information devs
}
