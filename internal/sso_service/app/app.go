package ssapp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"social-network/internal/sso_service/app/config"
	"social-network/internal/sso_service/infrastructure/repository"
	"social-network/internal/sso_service/transport/http/rest/controller"
	"social-network/internal/sso_service/transport/http/rest/router"
	"social-network/internal/sso_service/transport/http/rest/server"
	"social-network/internal/sso_service/usecase"
	gconfig "social-network/pkg/config"
	"social-network/pkg/infras/storage/postgresql"
)

// errors
var (
	ErrCritical = errors.New("error, critical error")
)

// Run
func Run(ctx context.Context, gcfg *gconfig.Config, dieChan <-chan struct{}) error {
	// init local configuration
	lcfg, err := config.InitSSOSRVConfig()
	if err != nil {
		return err
	}

	// InitPostgreSQL
	postgres, err := postgresql.New().Connection(ctx, gcfg)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	slog.Debug("Created a new PostgreSQL client")

	// Check the PostgreSQL connection
	if err := postgres.Ping(ctx); err != nil {
		slog.Error(err.Error())
		return err
	}
	slog.Info("Successful connected to the PostgreSQL")

	// Run Migration via goose cli utils & lib
	if err := postgres.RunMigration(gcfg.MigrationFilesPath); err != nil {
		slog.Error(err.Error())
		return err
	}
	slog.Info("Successful completed migrations via goose")
	defer postgres.Close()

	// TODO: RabbitMQ Consumer, and Producer

	// implement Dependency Inversion use Dependency Injection
	repo := repository.New(postgres)
	usecase := usecase.New(repo)
	controller := controller.New(usecase)

	// create and init router
	router := router.New(controller)
	router.Init(ctx)
	slog.Info("Successful Initialized SSO Service router")

	// create a new http sever instance
	httpServer := server.New(ctx, gcfg, lcfg, router)

	// information about service running
	info := fmt.Sprintf("The %s@v%s is running on: %s:%s",
		lcfg.Name,
		lcfg.Version,
		lcfg.Host,
		lcfg.Port,
	)
	slog.Info(info)

	// run the http server into separate goroutine for non blocking
	// the main service goroutine
	go func() {
		if err := httpServer.Start(ctx); err != nil {
			slog.Error(err.Error())
			// return
			// handler error different way
		}
	}()

	// graceful shutdown
	// stop the server
	<-dieChan // when receive from the dieChan msg, starting shutdown service
	slog.Info("Shuting down SSO Service")
	defer func(ctx context.Context) {
		if err := httpServer.Shutdown(ctx); err != nil {
			slog.Error(err.Error())
			return
		}
	}(ctx)
	slog.Debug("Close the SSO Service http server")

	// close db connection
	defer postgres.Close()
	slog.Debug("Close the SSO Server database connection")

	return ErrCritical
}
