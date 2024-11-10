package server

import (
	"context"
	"fmt"
	"net/http"

	"social-network/internal/sso_service/app/config"
	"social-network/internal/sso_service/transport/http/rest/router"
	gconfig "social-network/pkg/config"
)

// sso_service http server port

// HTTPServer
type HTTPServer struct {
	httpServer *http.Server
}

// New is constructor of the HTTPServer, return *HTTPServer
func New(ctx context.Context, cfg *gconfig.Config, lcfg *config.Config, handler *router.Router) *HTTPServer {
	addr := fmt.Sprintf("%s:%s", lcfg.HTTP.Host, lcfg.HTTP.Port)
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:           addr,
			WriteTimeout:   cfg.HTTPServer.WriteTimeout,
			ReadTimeout:    cfg.HTTPServer.ReadTimeout,
			IdleTimeout:    cfg.HTTPServer.IdleTimeout,
			MaxHeaderBytes: cfg.HTTPServer.MaxMB, // max header size
			Handler:        handler.Mux,
		},
	}
}

// Start the http server
// If can't start the http server return error
func (s *HTTPServer) Start(ctx context.Context) error {
	return s.httpServer.ListenAndServe()
}

// Shutdown the http server, receive ctx, stop the server
// If can't stop, return error
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
