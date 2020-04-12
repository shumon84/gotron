package file

import (
	"os"
	"path/filepath"
)

// Exists return true when a file exists, false otherwise.
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// ChmodWalkFunc return change mode func to use filepath.Walk.
func ChmodWalkFunc(mode os.FileMode) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, mode)
	}
}
