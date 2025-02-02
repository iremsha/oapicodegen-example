package main

import (
	"context"
	"fmt"

	"github.com/iremsha/oapicodegen-example/internal/app"
	"github.com/iremsha/oapicodegen-example/internal/config"
	logger "github.com/iremsha/oapicodegen-example/internal/log"
)

const (
	ConfigPath = "config.yml"
)

func main() {
	ctx := context.Background()
	log := logger.New()

	cfg, err := config.Load(ConfigPath)
	if err != nil {
		log.Error(ctx, fmt.Sprintf("Failed to load config: %s", err))
	}

	log.Info(ctx, "Starting oapicodegen-example server..")

	app := app.New(cfg, log)

	if err = app.Run(); err != nil {
		log.Error(ctx, fmt.Sprintf("Failed to start application: %s", err))
	}
}
