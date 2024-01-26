package userinfo

import (
	"os"
	"os/user"
	"path/filepath"
)

func UserDesktop() string {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	homeDir := curUser.HomeDir

	desktopPathRegular := filepath.Join(homeDir, "Desktop")
	desktopPathWithOneDrive := filepath.Join(homeDir, "OneDrive", "Desktop")

	desktopPath := desktopPathRegular

	if _, err := os.Stat(desktopPathRegular); os.IsNotExist(err) {
		// If desktopPathRegular doesn't exist, check if desktopPathWithOneDrive exists
		if _, err := os.Stat(desktopPathWithOneDrive); !os.IsNotExist(err) {
			desktopPath = desktopPathWithOneDrive
		}
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

	if _, err := os.Stat(downloadsPathRegular); os.IsNotExist(err) {
		// If downloadsPathRegular doesn't exist, check if downloadsPathWithOneDrive exists
		if _, err := os.Stat(downloadsPathWithOneDrive); !os.IsNotExist(err) {
			downloadsPath = downloadsPathWithOneDrive
		}
	}

	return downloadsPath
}
