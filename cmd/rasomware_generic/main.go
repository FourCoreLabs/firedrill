package main

import (
	"context"

	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransomware_generic"
	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		ransomware_generic.NewRansomEncrypt(),
		ransomware_generic.NewRansomNote(),
		ransomware_generic.NewRansomWallpaper(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
