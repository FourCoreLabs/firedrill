package discoveryprocess

import (
	"context"
	"fmt"
	"strings"

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
	logger.Sugar().Infof("Fetching information of running processes")

	processes, err := process.Processes()
	if err != nil {
		logger.Sugar().Warnf("error during process discovery: ", err.Error())
	}

	logger.Sugar().Infof("Dumping running processes")

	for _, proc := range processes {
		pid := proc.Pid           //ProcessID
		pname, err := proc.Name() //ProcessName

		if err != nil {
			logger.Sugar().Warnf("failed to fetch process name for pid %d: %s", pid, err.Error())
			continue
		}

		username, _ := proc.Username()

		if strings.Contains(fmt.Sprintf("%+v", err), "Access is denied.") {
			username = "SYSTEM"
		}

		logger.Sugar().Infow("Process Information", "pid", pid, "process_name", pname, "process_username", username)
	}

	return nil
}
