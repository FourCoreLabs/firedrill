package fileutils

import (
	"io"
	"os"
	"path/filepath"
)

func ReadTenFilesFromDirectory(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files := make([]string, 0, 10)
	for len(files) < 10 {
		infos, err := f.Readdir(10 - len(files))
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		for _, info := range infos {
			if info.Mode().IsRegular() {
				files = append(files, filepath.Join(dir, info.Name()))
			}
		}
	}
	return files, nil
}

func CopyFilesToTestFolder(filePaths []string, destFolder string) error {
	for _, filePath := range filePaths {
		sourceFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFilePath := filepath.Join(destFolder, filepath.Base(filePath))
		destFile, err := os.Create(destFilePath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		if err != nil {
			return err
		}
	}

	return nil
}
