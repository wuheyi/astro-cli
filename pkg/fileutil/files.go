package fileutil

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Exists returns a boolean indicating if the given path already exists
func Exists(path string) (bool, error) {
	if path == "" {
		return false, nil
	}

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if !os.IsNotExist(err) {
		return false, errors.Wrap(err, "cannot determine if path exists, error ambiguous")
	}

	return false, nil
}

// WriteStringToFile write a string to a file
func WriteStringToFile(path string, s string) error {
	return WriteToFile(path, strings.NewReader(s))
}

// WriteToFile writes an io.Reader to a file if it does not exst
func WriteToFile(path string, r io.Reader) error {
	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}
