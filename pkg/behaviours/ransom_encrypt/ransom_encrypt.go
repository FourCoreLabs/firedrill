package ransom_encrypt

import (
	"context"
	"embed"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"go.uber.org/zap"
)

const (
	ID   = "ransom_encrypt"
	Name = "Ransomware Encryption"

	ransomDirName  = "fireDrillRansomware"
	embedFilesPath = "testfiles"
)

// files are test file to be dropped on the file system and encrypted as part of ransomware encryption simulation.
//go:embed testfiles
var files embed.FS

type RansomEncryptOptions struct {
	RansomwareDirName string
}

type RansomEncrypt struct {
	ransomDirName string
}

func NewRansomEncrypt(opts ...RansomEncryptOptions) sergeant.Runnable {
	var options RansomEncryptOptions = RansomEncryptOptions{
		RansomwareDirName: ransomDirName,
	}
	if len(opts) > 0 {
		options = opts[0]
	}

	return &RansomEncrypt{ransomDirName: options.RansomwareDirName}
}

func (e *RansomEncrypt) ID() string {
	return ID
}

func (e *RansomEncrypt) Name() string {
	return Name
}

func (e *RansomEncrypt) Run(ctx context.Context, logger *zap.Logger) error {
	desktopPath := UserDesktop()
	logger.Sugar().Infof("User desktop path for ransomare encryption: %s", desktopPath)

	testFiles, _ := files.ReadDir(embedFilesPath)

	targetDirPath := filepath.Join(desktopPath, e.ransomDirName)

	if _, err := os.Stat(targetDirPath); os.IsExist(err) {
		if err := os.RemoveAll(targetDirPath); err != nil {
			logger.Sugar().Warnf("Failed to delete old test folder %s: %s", targetDirPath, err.Error())
		}
	}

	if err := os.Mkdir(targetDirPath, 0755); err != nil {
		logger.Sugar().Warnf("Failed to make target folder for ransomware %s: %s", targetDirPath, err.Error())
		return err
	}
	logger.Sugar().Infof("Generated test folder for ransomware encryption: %s", targetDirPath)

	for _, file := range testFiles {
		targetFilePath := filepath.Join(targetDirPath, file.Name())
		testFilePath := path.Join(embedFilesPath, file.Name())

		testFileBuf, err := files.ReadFile(testFilePath)
		if err != nil {
			return err
		}

		if err := os.WriteFile(targetFilePath, testFileBuf, 0644); err != nil {
			return err
		}
	}

	return nil
}

func UserDesktop() string {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	homeDir := curUser.HomeDir
	return filepath.Join(homeDir, "Desktop")
}
