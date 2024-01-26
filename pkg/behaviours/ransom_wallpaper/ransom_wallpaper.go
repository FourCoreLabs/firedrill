package ransom_wallpaper

import (
	"context"
	_ "embed"
	"fmt"
	"runtime"
	"time"

	"github.com/FourCoreLabs/firedrill/pkg/sergeant"
	"github.com/FourCoreLabs/firedrill/pkg/utils/winwallpaper"
	"go.uber.org/zap"
)

//go:embed ransom.jpg
var ransomWallpaperBuf []byte

const (
	ID   = "ransom_wallpaper"
	Name = "Ransomware Wallpaper"
)

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
	// embeddedBufferLength := len(ransomWallpaperBuf)
	var options RansomWallpaperOptions = RansomWallpaperOptions{
		CurrentWallpaperPath: wallpaperPath,
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	return &RansomWallpaper{currentWallpaperPath: options.CurrentWallpaperPath, embeddedWallpaperLength: len(ransomWallpaperBuf)}
}

func (e *RansomWallpaper) ID() string {
	return ID
}

func (e *RansomWallpaper) Name() string {
	return Name
}

func (e *RansomWallpaper) Run(ctx context.Context, logger *zap.Logger) error {

	logger.Sugar().Infof("Current Wallpaper Path: %s", e.currentWallpaperPath)
	logger.Sugar().Infof("Embedded Ransom Wallpaper size: %d", e.embeddedWallpaperLength)

	switch runtime.GOOS {
	case "windows":
		wallpaperErr := winwallpaper.ChangeSystemWallpaperwithConfig(ransomWallpaperBuf, e.currentWallpaperPath)
		if wallpaperErr != nil {
			logger.Sugar().Warnf(fmt.Sprintf("error during wallpaper change: %v", wallpaperErr))
		}
		logger.Sugar().Infof("Changed system wallpaper, original wallpaper path backed up in registry")

		logger.Sugar().Infof("Sleeping for 07 seconds")
		time.Sleep(07 * time.Second)

		wallpaperErr = winwallpaper.RestoreWallPaperfromConfig()
		if wallpaperErr != nil {
			logger.Sugar().Warnf(fmt.Sprintf("error during wallpaper restore: %v", wallpaperErr))
		}
		logger.Sugar().Infof("Restored existing wallpaper")
	default:
	}
	return nil
}
