package discoveryprocess

import (
	"context"
	"errors"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"github.com/shirou/gopsutil/v3/process"
	"go.uber.org/zap"
)

const (
	ID   = "discovery_process"
	Name = "Discovery Process"
)

type DiscoveryProcessOptions struct{}

type DiscoveryProcess struct{}

func NewDiscoveryProcess(opts ...DiscoveryProcessOptions) sergeant.Runnable {
	// var options DiscoveryProcessOptions = DiscoveryProcessOptions{}

	// if len(opts) > 0 {
	// 	options = opts[0]
	// }

	return &DiscoveryProcess{}
}

func (e *DiscoveryProcess) ID() string {
	return ID
}

func (e *DiscoveryProcess) Name() string {
	return Name
}

func (e *DiscoveryProcess) Run(ctx context.Context, logger *zap.Logger) error {
	logger.Sugar().Infof("T1057 Process Discovery: Fetching information of running processes")

	processes, err := process.Processes()
	if err != nil {
		logger.Sugar().Warnf("error during process discovery: ", err.Error())
	}

	logger.Sugar().Infof("Dumping running processes")

	for _, proc := range processes {
		pid := proc.Pid           //ProcessID
		pname, err := proc.Name() //ProcessName

		if err != nil {
			logger.Sugar().Warnf("failed to fetch process name for pid %d/%s: %s", pid, pname, err.Error())
			continue
		}

		username, err := proc.Username()
		if err != nil {
			if errors.Is(err, errors.New("Access is denied.")) {
				username = "SYSTEM"
			} else {
				logger.Sugar().Warnf("failed to fetch username for pid %d/%s: %s", pid, pname, err.Error())
			}
		}

		logger.Sugar().Infow("Process Information", "pid", pid, "process_name", pname, "process_username", username)
	}

	return nil
}
