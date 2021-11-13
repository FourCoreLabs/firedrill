package echo

import (
	"context"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

const (
	ID   = "echo"
	Name = "Echo behaviour"
)

type EchoOptions struct {
	Message string
}

type Echo struct {
	message string
}

func NewEcho(opts ...EchoOptions) sergeant.Runnable {
	var options EchoOptions = EchoOptions{
		Message: "Hello World",
	}
	if len(opts) > 0 {
		options = opts[0]
	}

	return &Echo{message: options.Message}
}

func (e *Echo) ID() string {
	return ID
}

func (e *Echo) Name() string {
	return Name
}

func (e *Echo) Run(ctx context.Context, logger *zap.Logger) error {
	logger.Sugar().Infof("echo: %s", e.message)

	return nil
}
