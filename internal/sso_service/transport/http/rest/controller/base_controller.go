package controller

import (
	"fmt"
	"net/http"
)

// SSOUseCase
type SSOUseCase interface {
	// TODO: SSO Service use cases
}

// Type Controller
type Controller struct {
	SSOUseCase // Dependency injection
}

// New is constructor of the Controller
// return a new *Controller instance
func New(su SSOUseCase) *Controller {
	return &Controller{
		SSOUseCase: su,
	}
}

// Auth

// User

// Profile

func (c *Controller) BaseController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("\nbase controller\n"))

	fmt.Println("Base controller accepted request and answer: base controller")
}
