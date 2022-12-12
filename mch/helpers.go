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

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func IfFn[T any](fn func() bool, a, b T) T {
	if fn() {
		return a
	}
	return b
}
