// Package file implements filesystem I/O functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Create creates a new file populated from a string.
func Create(dest, body string, mode os.FileMode) error {
	if Exists(dest) {
		return fmt.Errorf("cannot create file %q - already exists", dest)
	}

	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot create file %q - %w", dest, err)
	}

	return nil
}

// Delete renames a file to the ".trash" extension.
func Delete(orig string) error {
	if !Exists(orig) {
		return fmt.Errorf("cannot delete file %q - does not exist", orig)
	}

	extn := filepath.Ext(orig)
	dest := strings.Replace(orig, extn, ".trash", 1)
	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot delete file %q - %w", orig, err)
	}

	return nil
}

// Exists returns true if a directory or file exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// Read returns a file's body as a string.
func Read(orig string) (string, error) {
	if !Exists(orig) {
		return "", fmt.Errorf("cannot read file %q - does not exist", orig)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return "", fmt.Errorf("cannot read file %q - %w", orig, err)
	}

	return string(bytes), nil
}

// Update overwrites an existing file with a string.
func Update(orig, body string, mode os.FileMode) error {
	if !Exists(orig) {
		return fmt.Errorf("cannot update file %q - does not exist", orig)
	}

	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot update file %q - %w", orig, err)
	}

	return nil
}
