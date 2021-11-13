package main

import (
	"context"

	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransom_encrypt"
	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransom_note"
	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		ransom_encrypt.NewRansomEncrypt(),
		ransom_note.NewRansomNote(),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
