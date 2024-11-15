package ssapp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"social-network/internal/platform/infras"
	"social-network/internal/sso_service/app/config"
	ssorepository "social-network/internal/sso_service/infrastructure/repository"
	"social-network/internal/sso_service/infrastructure/repository/storage/postgres"
	grpcserver "social-network/internal/sso_service/transport/grpc/server"
	"social-network/internal/sso_service/transport/http/rest/controller"
	"social-network/internal/sso_service/transport/http/rest/router"
	"social-network/internal/sso_service/transport/http/rest/server"
	ssousecase "social-network/internal/sso_service/usecase"
	gconfig "social-network/pkg/config"
)

// errors
var (
	ErrCritical = errors.New("error, critical error")
)

// Run
func Run(
	ctx context.Context,
	gcfg *gconfig.Config,
	infra *infras.Infras,
	dieChan <-chan struct{},
	crtErrChan chan<- error,
) {
	// init local configuration
	lcfg, err := config.InitSSOSRVConfig()
	if err != nil {
		crtErrChan <- err
		return
	}

	// TODO: RabbitMQ Consumer, and Producer
	// TODO: Redis, cache aside strategy, with the T-LRU alg of the caching

	// NewPostgreSQL client
	postgres := postgres.New(infra.PostgreSQL)

	// implement Dependency Inversion use Dependency Injection
	repo := ssorepository.New(postgres)
	usecase := ssousecase.New(repo)
	controller := controller.New(usecase)

	// create and init router
	router := router.New(controller)
	router.Init(ctx)
	slog.Info("Successful Initialized SSO Service router")

	// create a new http sever instance

	httpServer := server.New(ctx, gcfg, lcfg, router)
	fmt.Println("HTTPServer ", httpServer == nil)

	// information about service running
	info := fmt.Sprintf("The %s@v%s is running on: %s:%s",
		lcfg.Name,
		lcfg.Version,
		lcfg.HTTP.Host,
		lcfg.HTTP.Port,
	)
	slog.Info(info)

	// new grpc server

	grpcServer, err := grpcserver.New(ctx, gcfg, lcfg, usecase, slog.Default())
	if err != nil {
		slog.Error(err.Error())
		crtErrChan <- err
		return
	}

	// information about grpc server starting
	ginfo := fmt.Sprintf("The %s@v%s is running on: %s:%s",
		lcfg.Name,
		lcfg.Version,
		lcfg.GRPC.Host,
		lcfg.GRPC.Port,
	)
	slog.Info(ginfo)
	// run gRPC Server
	go func() {
		if err := grpcServer.Run(); err != nil {
			slog.Error(err.Error())
		}
	}()

	// run the http server into separate goroutine for non blocking
	// the main service goroutine
	go func() {
		if err := httpServer.Start(ctx); err != nil {
			slog.Error(err.Error())
			// handler error different way
		}
	}()

	// graceful shutdown
	// stop the server
	<-dieChan // when receive from the dieChan msg, starting shutdown service
	slog.Info("Shuting down SSO Service")
	if err := httpServer.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Debug("Close the SSO Service http server")

	crtErrChan <- ErrCritical
}
