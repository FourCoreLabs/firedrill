package ransom_wallpaper

import (
	"context"
	_ "embed"
	"fmt"
	"runtime"

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
		wallpaperErr := winwallpaper.ChangeSystemWallpaper(ransomWallpaperBuf)
		if wallpaperErr != nil {
			logger.Sugar().Warnf(fmt.Sprintf("error during wallpaper change: %v", wallpaperErr))
		}
		logger.Sugar().Infof("Changed system wallpaper")
	default:
	}
	return nil
}
