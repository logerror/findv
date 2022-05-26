package utils

import (
	"os"
	"path/filepath"
)

var cacheDir string

func DefaultCacheDir() string {
	tmpDir, err := os.UserCacheDir()
	if err != nil {
		tmpDir = os.TempDir()
	}
	return filepath.Join(tmpDir, "findv")
}

func SetCacheDir(dir string) {
	cacheDir = dir
}
