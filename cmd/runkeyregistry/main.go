package main

import (
	"context"

	"github.com/FourCoreLabs/firedrill/pkg/behaviours/runkey_registry"
	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		runkey_registry.NewRegistryRunKey(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
