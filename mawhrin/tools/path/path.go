// Package path implements file path manipulation functions.
package path

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

// Glob returns a sorted slice of all paths in a directory matching an extension.
func Glob(dire, extn string) ([]string, error) {
	glob := filepath.Join(dire, "*"+extn)
	paths, err := filepath.Glob(glob)
	if err != nil {
		return nil, fmt.Errorf("cannot list directory %q", dire)
	}

	slices.Sort(paths)
	return paths, nil
}

// Join returns a joined path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Name returns a path's base name without the extension.
func Name(path string) string {
	base := filepath.Base(path)
	return strings.SplitN(base, ".", 2)[0]
}
