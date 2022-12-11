package mch

import (
	"errors"
	"io/fs"
	"os"
)

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}
