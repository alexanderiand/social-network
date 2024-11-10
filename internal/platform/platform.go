package platform

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"social-network/internal/platform/infras"
	ssoapp "social-network/internal/sso_service/app"
	"social-network/pkg/config"
)

var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// Run Platform with of the platform microservices, receive ctx, and *cfg
// If a *Config struct pointer is nil, return ErrNilStructPointer
// If the *Config params is invalid, return ErrInvalidCfgParam
// If error accrued while initializing the databases new instance, return ErrDBInit*(postgres or other db error)
// If error happened while implement DI, return more specific Err*(layer init error)
// If going unknown error, return this unknown error
func Run(
	ctx context.Context,
	cfg *config.Config,
	infra *infras.Infras,
	crtErrChan chan error,
	dieChan chan struct{},
) {

	//! every microservices implement into itself Dependency Inversion, use Dependency Injection

	// srv - service
	// TODO: eventsrv.Run

	// ssosrv.Run
	if err := ssoapp.Run(ctx, cfg, infra, dieChan); err != nil {
		// call platform.Stop
		crtErrChan <- err
		return
	}

	// TODO: contentsrv.Run

	// TODO: chatsrv.Run

	// TODO: comsrv.Run
}

// Stop called if context canceled, receive os.Signal, or critical error, or invalid config params
// Stop implement Graceful Shutdown for the Platform, also for every sub services of the Platform
// If ctx, sigChan, or crtErrChan is invalid, return Err*(above arg name)
// If the databases closing with error, return this specific error as a wrapped error with more info
// If the message broker, RabbitMQ return error in closing, return this error as a wrapped error with more error
// if another going wrong with unknown error, return this error as the wrapped error, with additional information
func Stop(
	ctx context.Context,
	cfg *config.Config,
	critErrChan chan error,
	dieChan chan struct{},
	wg *sync.WaitGroup,
	srvCount int,
) {
	// receive context canceling, os signal or critical error, start graceful shutdown the Platform
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

loop:
	for {
		select {
		case <-ctx.Done():
			slog.Error("Parent context canceled! Starting shutdown the platform")
			break loop
		case err := <-critErrChan:
			if err != nil {
				critErrInfo := "Critical error: "
				slog.Error(fmt.Errorf("%s %w", critErrInfo, err).Error())
				critErrChan <- err
				break loop
			}
			continue
		case sig := <-sigChan:
			info := fmt.Sprintf(
				"Receive the os.Signal: %s, starting shuting down the platform",
				sig.String())
			slog.Error(info)
			break loop
		}
	}

	// starting shutdown the platform
	genDieSignal(srvCount, dieChan)

	wg.Done()
}

// genDieSignal
func genDieSignal(srvCount int, dieChan chan struct{}) {
	for srvCount > 0 {
		dieChan <- struct{}{}
		srvCount--
	}
}
