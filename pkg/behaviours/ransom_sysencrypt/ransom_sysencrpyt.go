package ransom_sysencrypt

import (
	"context"
	_ "embed"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"github.com/FourCoreLabs/firedrill/pkg/utils/aesutils"
	"github.com/FourCoreLabs/firedrill/pkg/utils/fileutils"
	"github.com/FourCoreLabs/firedrill/pkg/utils/userinfo"
	"go.uber.org/zap"
)

const (
	ID            = "ransom_sysencrypt"
	Name          = "Ransomware Local FilesEncryption"
	ext           = ".drill"
	ransomDirName = "fireDrillRansomware"
)

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
	desktopPath := userinfo.UserDesktop()
	downloadsPath := userinfo.UserDownloads()
	targetDirPath := filepath.Join(desktopPath, e.ransomDirName)

	logger.Sugar().Infof("User Desktop and Downloads path for ransomware encryption: %s, %s", desktopPath, downloadsPath)
	logger.Sugar().Infof("Creating target directory for ransomware simulation")

	if _, err := os.Stat(targetDirPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(targetDirPath); err != nil {
			logger.Sugar().Warnf("Failed to delete old test folder %s: %s", targetDirPath, err.Error())
		}
	}

	if err := os.Mkdir(targetDirPath, 0755); err != nil {
		logger.Sugar().Warnf("Failed to make target folder for ransomware %s: %s", targetDirPath, err.Error())
		return err
	}

	logger.Sugar().Infof("Copying User files from Desktop and Dowloads folder to test directory for ransomare encryption")
	desktopfilePaths, err := fileutils.ReadTenFilesFromDirectory(desktopPath)
	if err != nil {
		logger.Sugar().Warnf("Failed to read files from Desktop: %s", err.Error())
	}

	downloadfilePaths, err := fileutils.ReadTenFilesFromDirectory(downloadsPath)
	if err != nil {
		logger.Sugar().Warnf("Failed to read files from Downloads: %s", err.Error())
	}

	collectedfilePaths := append(desktopfilePaths, downloadfilePaths...)
	if len(collectedfilePaths) == 0 {
		logger.Sugar().Fatal("no files collected for ransomware simulation")
	}

	err = fileutils.CopyFilesToTestFolder(collectedfilePaths, targetDirPath)
	if err != nil {
		logger.Sugar().Fatalf("Failed to copy files in the target directory: %s", err.Error())
	}

	logger.Sugar().Infof("Created test folder with user folder: %s", targetDirPath)

	aesKey := aesutils.AESEncryptionKey()
	files, err := os.ReadDir(targetDirPath)
	if err != nil {
		return err
	}

	totalFiles := len(files)

	logger.Sugar().Infof("Encrypting %d files.", totalFiles)

	for i, file := range files {
		fileAbsPath := filepath.Join(targetDirPath, file.Name())
		fileData, err := os.ReadFile(fileAbsPath)
		if err != nil {
			return err
		}

		encData, err := aesutils.AESEncryptData(fileData, aesKey)
		if err != nil {
			return err
		}

		encFilePath := fileAbsPath + ext

		if err := os.WriteFile(encFilePath, encData, 0644); err != nil {
			return err
		}

		if err := os.Remove(fileAbsPath); err != nil {
			return err
		}

		logger.Sugar().Infof("Encrypted %d/%d files.", i+1, totalFiles)
	}

	logger.Sugar().Info("Waiting for 03 seconds.")

	select {
	case <-time.After(3 * time.Second):
	case <-ctx.Done():
		return errors.New("context cancelled")
	}

	return nil
}
