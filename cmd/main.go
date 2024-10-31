package main

import (
	"log"
	"log/slog"
	"social-network/pkg/config"
)

func main() {
	// TODO: init config, error haanling
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Failed to init config", err) // Пока так оставил в будущем реализуем норм логер.
	}
	slog.Debug("Successful project initialization\n Service name:%v, Service version:%v \n", cfg.Service.Name, cfg.Service.Version)
}
