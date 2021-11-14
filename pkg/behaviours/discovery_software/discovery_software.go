package discoverysoftware

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"sync"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

const (
	ID         = "discovery_software"
	Name       = "Discovery Software"
	ShellToUse = "powershell.exe"
)

type DiscoverySoftwareOptions struct{}

type DiscoverySoftware struct{}

func NewDiscoverySoftware(opts ...DiscoverySoftwareOptions) sergeant.Runnable {
	return &DiscoverySoftware{}
}

func (e *DiscoverySoftware) ID() string {
	return ID
}

func (e *DiscoverySoftware) Name() string {
	return Name
}

func (e *DiscoverySoftware) Run(ctx context.Context, logger *zap.Logger) error {
	logger.Sugar().Infof("T1518 Installed Software Discovery: Fetching a list of installed software")
	var params, output []string
	var paramWg sync.WaitGroup
	var outputLock sync.Mutex
	params = append(params, `Get-ItemProperty HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName, DisplayVersion, Publisher, InstallDate | Format-Table -Autosize`, `Get-ItemProperty HKLM:\Software\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName, DisplayVersion, Publisher, InstallDate | Format-Table -Autosize`)
	for x := range params {
		paramWg.Add(1)
		go func(args string) {
			defer paramWg.Done()
			stdout, _, err := ExecShell(args)

			if err != nil {
				return
			}
			outputLock.Lock()
			defer outputLock.Unlock()
			if len(stdout) != 0 {
				output = append(output, string(stdout))
			}
		}(params[x])
	}
	paramWg.Wait()
	if len(output) != 0 {
		logger.Sugar().Infof("Printing list of installed software: \n")
		fmt.Print(output)
	}
	return nil
}

func ExecShell(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	execErr := cmd.Run()
	return stdout.String(), stderr.String(), execErr
}
