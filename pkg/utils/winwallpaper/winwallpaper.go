package winwallpaper

import (
	"errors"
	"os"

	"github.com/FourCoreLabs/firedrill/pkg/utils/regutils"
	"github.com/reujab/wallpaper"
)

// This function returns the path of the current wallpaper.
func GetCurrentWallpaperPath() (string, error) {
	currentPath, err := wallpaper.Get()
	if err != nil {
		return "", err
	}
	return currentPath, nil
}

// This function takes a byte array as an image to be used to replace the current wallpaper.
func ChangeSystemWallpaper(wallpaperbuf []byte) error {
	tempfile, err := os.CreateTemp(os.TempDir(), "fctemp*.jpg")
	if err != nil {
		return err
	}
	_, err = tempfile.Write(wallpaperbuf)
	if err != nil {
		return err
	}
	filepath := tempfile.Name()
	tempfile.Close()
	err = wallpaper.SetFromFile(filepath)
	if err != nil {
		return err
	}
	return nil
}

func ChangeSystemWallpaperwithConfig(wallpaperbuf []byte, currentPath string) error {

	// save current path to registry
	if err := regutils.SaveConfigToRegistry("WallpaperBackup", currentPath); err != nil {
		return err
	}

	if err := ChangeSystemWallpaper(wallpaperbuf); err != nil {
		return err
	}

	return nil
}

func RestoreWallPaperfromConfig() error {

	key, err := regutils.GetConfigFromRegistry("WallpaperBackup")
	if err != nil {
		return err
	}
	path, ok := key.(string)
	if !ok {
		return errors.New("unable to get config string from interface")
	}
	err = wallpaper.SetFromFile(path)
	if err != nil {
		return err
	}
	return nil
}
