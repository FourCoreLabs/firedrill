package main

import (
	"context"

	"github.com/FourCoreLabs/firedrill/pkg/behaviours/echo"
	"github.com/FourCoreLabs/firedrill/pkg/behaviours/ransom_note"
	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	behaviours := []sergeant.Runnable{
		ransom_note.NewRansomNote(),
		echo.NewEcho(echo.EchoOptions{Message: "this is ransomware, hehe"}),
	}

	sergeant := sergeant.NewSergeant(logger, behaviours...)
	if err := sergeant.Start(context.Background()); err != nil {
		logger.Sugar().Fatalw("execution failed", "error", err.Error())

	}
}
