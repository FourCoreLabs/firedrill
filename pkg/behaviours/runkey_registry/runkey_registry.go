package runkey_registry

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

//Embed Persistence payload
//go:embed persistence_blob.exe
var pBlob []byte
var pFile string

func init() {
	tempfile, err := os.CreateTemp(os.TempDir(), "*.exe")
	if err != nil {
		panic(err)
	}
	_, err = tempfile.Write(pBlob)
	if err != nil {
		panic(err)
	}
	tempfile.Close()
	pFile = tempfile.Name()
}

const (
	ID   = "registryrunkey"
	Name = "Registry Run Key Persistence"
)

type RegistryRunKeyOptions struct {
	path string
}

type RegistryRunKey struct {
	Path string
}

func NewRegistryRunKey(opts ...RegistryRunKeyOptions) sergeant.Runnable {
	var options RegistryRunKeyOptions = RegistryRunKeyOptions{
		path: pFile,
	}
	if len(opts) > 0 {
		options = opts[0]
	}

	return &RegistryRunKey{Path: options.path}
}

func (e *RegistryRunKey) ID() string {
	return ID
}

func (e *RegistryRunKey) Name() string {
	return Name
}

func (e *RegistryRunKey) Run(ctx context.Context, logger *zap.Logger) error {
	sugared := logger.Sugar()

	sugared.Info("Performing Registry Run Key Persistence")

	// Setting up the registry format to create the bypass
	command := fmt.Sprintf(`REG ADD "HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /V "FIREDRILL" /t REG_SZ /F /D "%s"`, e.Path)

	sugared.Infof("Command to be executed: %s", command)

	_, err := exec.Command("C:\\Windows\\System32\\cmd.exe", "/C", command).Output()
	if err != nil {
		return err
	}

	sugared.Infof("Payload path persisted: %s", e.Path)

	sugared.Info("Resetting registry settings to remove Run Key")

	command = `REG DELETE "HKCU\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /V "FIREDRILL" /f`
	sugared.Infof("Command to be executed: %s", command)

	_, err = exec.Command("C:\\Windows\\System32\\cmd.exe", "/C", command).Output()
	if err != nil {
		return err
	}
	sugared.Info("Deleted Registry Key")
	return nil
}
