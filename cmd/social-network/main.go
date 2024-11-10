package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"

	"social-network/internal/platform"
	"social-network/pkg/config"
	"social-network/pkg/logger"
)

// _srtCount total service count, need for implementing shuting down platform
const _srvCount = 1

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Printf("here: %v", err)
		log.Fatal("Failed to init config\n", err)
	}
	log.Printf("Successful project initialization Service name:%v, Service version:%v", cfg.Service.Name, cfg.Service.Version)

	if _, err := logger.InitLogger(cfg); err != nil {
		log.Fatal("Failed to init logger\n", cfg)
	}
	slog.Debug("Successful project initialization Service name:%v, Service version:%v", cfg.Service.Name, cfg.Service.Version)

	// Parent Context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// platform.Stop args
	dieChan := make(chan struct{}, _srvCount)
	crtErrChan := make(chan error, 2)
	defer close(crtErrChan)
	defer close(dieChan)
	wg := &sync.WaitGroup{}

	// initInfrastructure
	infra, err := platform.InitInfrastructures(ctx, cfg)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Infrastructures successful initialized")

	// run platform
	go platform.Run(ctx, cfg, infra, crtErrChan, dieChan)

	// stop platform
	wg.Add(1)
	go platform.Stop(ctx, cfg, crtErrChan, dieChan, wg, _srvCount)

	// if receive critical error, cancel parent context
	if err := <-crtErrChan; err != nil {
		slog.Error(err.Error())
		cancel()
	}

	// waiting goroutines

	wg.Wait()

	// close infras services connections
	if err := platform.CloseConnections(ctx, infra); err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Closed the Infrastructure services connections")

	slog.Info("Bye! The Platform successful stopped!")
}
