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

	// TODO: run platform
	if err := platform.Run(ctx, cfg); err != nil {
		slog.Error(err.Error())
		slog.Info("Parent context canceled!")
		cancel()
	}

	// TODO: stop platform
}
