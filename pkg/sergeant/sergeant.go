package sergeant

import (
	"context"
	"errors"
	"sync"

	"go.uber.org/zap"
)

type Runnable interface {
	ID() string
	Name() string
	Run(ctx context.Context, l *zap.Logger) error
}

type Sergeant struct {
	runnables   []Runnable
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	runnableMu  sync.Mutex
}

func (s *Sergeant) AddRunnable(runnable Runnable) {
	s.runnableMu.Lock()
	defer s.runnableMu.Unlock()

	s.runnables = append(s.runnables, runnable)
}

func (s *Sergeant) Start(ctx context.Context) error {
	for _, runnable := range s.runnables {
		select {
		case <-ctx.Done():
			return errors.New("execution context cancelled")
		default:
		}

		s.sugarLogger.Infow("Starting execution", "id", runnable.ID(), "name", runnable.Name())

		if err := runnable.Run(ctx, s.logger); err != nil {
			return err
		}

		s.sugarLogger.Infow("Finished execution", "id", runnable.ID(), "name", runnable.Name())
	}

	return nil
}

func NewSergeant(logger *zap.Logger, runnables ...Runnable) *Sergeant {
	sergeant := &Sergeant{
		runnables:   runnables,
		logger:      logger,
		sugarLogger: logger.Sugar(),
	}

	return sergeant
}
