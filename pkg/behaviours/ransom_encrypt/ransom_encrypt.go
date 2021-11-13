package ransom_encrypt

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"embed"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"github.com/FourCoreLabs/firedrill/pkg/utils/userinfo"
	"go.uber.org/zap"
)

const (
	ID   = "ransom_encrypt"
	Name = "Ransomware Encryption"

	ext            = ".drill"
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
	desktopPath := userinfo.UserDesktop()
	logger.Sugar().Infof("User desktop path for ransomare encryption: %s", desktopPath)

	testFiles, _ := files.ReadDir(embedFilesPath)

	targetDirPath := filepath.Join(desktopPath, e.ransomDirName)

	if _, err := os.Stat(targetDirPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(targetDirPath); err != nil {
			logger.Sugar().Warnf("Failed to delete old test folder %s: %s", targetDirPath, err.Error())
		}
	}

	if err := os.Mkdir(targetDirPath, 0755); err != nil {
		logger.Sugar().Warnf("Failed to make target folder for ransomware %s: %s", targetDirPath, err.Error())
		return err
	}

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
	logger.Sugar().Infof("Generated test folder for ransomware encryption: %s", targetDirPath)

	aesKey := aesEncryptionKey()

	files, err := os.ReadDir(targetDirPath)
	if err != nil {
		return err
	}

	totalFiles := len(files)
	encFilePaths := make([]string, 0, totalFiles)

	logger.Sugar().Infof("Encrypting %d files.", totalFiles)

	for i, file := range files {
		fileAbsPath := filepath.Join(targetDirPath, file.Name())
		fileData, err := os.ReadFile(fileAbsPath)
		if err != nil {
			return err // everything should work.
		}

		encData, err := aesEncryptData(fileData, aesKey)
		if err != nil {
			return err // everything should work.
		}

		encFilePath := fileAbsPath + ext

		if err := os.WriteFile(encFilePath, encData, 0644); err != nil {
			return err
		}

		if err := os.Remove(fileAbsPath); err != nil {
			return err
		}

		encFilePaths = append(encFilePaths, encFilePath)
		logger.Sugar().Infof("Encrypted %d/%d files.", i+1, totalFiles)
	}

	logger.Sugar().Info("Waiting for 10 seconds.")

	select {
	case <-time.After(10 * time.Second):
	case <-ctx.Done():
		return errors.New("context cancelled")
	}

	return nil
}

// aesEncryptData encrypts data using 256-bit AES-GCM. Output: nonce+cipherdata+tag
func aesEncryptData(data []byte, key []byte) (encryptedtext []byte, err error) {
	cipherblock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcmpack, err := cipher.NewGCM(cipherblock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcmpack.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcmpack.Seal(nonce, nonce, data, nil), nil
}

// aesEncryptionKey returns random AES Encrpytion Key
func aesEncryptionKey() []byte {
	ekey := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, ekey)
	if err != nil {
		panic(fmt.Sprintf("Failed to seed key: %v", err))
	}
	return ekey
}
