package discoveryperipheral

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

const (
	ID         = "discovery_peripheral"
	Name       = "Discovery Peripheral"
	ShellToUse = "powershell.exe"
)

type DiscoveryPeripheralOptions struct{}

type DiscoveryPeripheral struct{}

func NewDiscoveryPeripheral(opts ...DiscoveryPeripheralOptions) sergeant.Runnable {
	return &DiscoveryPeripheral{}
}

func (e *DiscoveryPeripheral) ID() string {
	return ID
}

func (e *DiscoveryPeripheral) Name() string {
	return Name
}

func (e *DiscoveryPeripheral) Run(ctx context.Context, logger *zap.Logger) error {
	logger.Sugar().Infof("T1120 Peripheral Discovery: Fetching information of connected peripherals and drivers")
	logger.Sugar().Infof("Enumerating connected devices")
	devices, _, execErr := ExecShell(`pnputil /enum-devices /connected`)
	if execErr != nil {
		logger.Sugar().Warnf("error enumerating connected devices through pnputil: ", execErr)
	}
	fmt.Print(devices)
	// logger.Sugar().Infof("%s", devices)
	logger.Sugar().Infof("Enumerating installed drivers")
	drivers, _, execErr := ExecShell(`pnputil /enum-drivers`)
	if execErr != nil {
		logger.Sugar().Warnf("error enumerating drivers through pnputil: ", execErr)
	}
	fmt.Print(drivers)
	// logger.Sugar().Infof(fmt.Sprintf("%v", drivers))
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
