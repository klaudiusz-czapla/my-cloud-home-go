package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Determines whether the specified file exists.
func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

// Opens a text file (if exists), reads all the text in the file into a string, and then closes the file.
func ReadAllText(f string) (string, error) {
	if !FileExists(f) {
		return "", fmt.Errorf("file %s does not exist", f)
	}
	data, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Creates a new file, write the contents to the file, and then closes the file. If the target file already exists, it is overwritten.
func WriteAllText(f, data string) error {
	file, err := os.OpenFile(f, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(int(0600)))
	if err != nil {
		return err
	}

	file.WriteString(data)

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

// Emulates coalesce operator known from other languages
func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

// Emulates coalesce operator known from other languages
func IfFn[T any](fn func() bool, a, b T) T {
	if fn() {
		return a
	}
	return b
}
