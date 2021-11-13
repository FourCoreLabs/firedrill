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

	if _, err := os.Stat(desktopPathWithOneDrive); !os.IsNotExist(err) {
		desktopPath = desktopPathWithOneDrive
	}

	return desktopPath
}
