package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"social-network/internal/platform"
	"social-network/pkg/config"
	"social-network/pkg/logger"
)

func main() {
	// init config, error haanling
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Printf("here: %v", err)
		log.Fatal("Failed to init config\n", err)
	}
	log.Printf("Successful project initialization Service name:%v, Service version:%v", cfg.Service.Name, cfg.Service.Version)
	
	// init logger  
	if _, err := logger.InitLogger(cfg); err != nil {
		log.Fatal("Failed to init logger\n", cfg)
	}

	slog.Debug("Successful project initialization Service name:%v, Service version:%v", cfg.Service.Name, cfg.Service.Version)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	if err := platform.Run(ctx, cfg); err != nil {
		slog.Error(err.Error())
		cancel()
	}

	// platfrom.Run
	// error handling
	// call platfrom.Stop
	// cancel context
}
