package main

import (
	"fmt"
	"log"
	"log/slog"
	"social-network/pkg/config"
	"social-network/pkg/logger"
)

func main() {
	// TODO: init config, error haanling
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Printf("here: %v", err)
		log.Fatal("Failed to init config\n", err)
	}
	log.Printf("Successful project initialization Service name:%v, Service version:%v", cfg.Service.Name, cfg.Service.Version)

	fmt.Printf("\nCFG: %+v\n", cfg)
	// TODO: init logger
	if _, err := logger.InitLogger(cfg); err != nil {
		log.Fatal("Failed to init logger\n", cfg)
	}

	slog.Debug("Successful project initialization Service name:%v, Service version:%v", cfg.Service.Name, cfg.Service.Version)
}
