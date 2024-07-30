// Package path implements file path manipulation functions.
package path

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a path's file extension with a leading dot.
func Extn(orig string) string {
	return filepath.Ext(orig)
}

// Glob returns a sorted slice of all paths in a directory matching an extension.
func Glob(dire, extn string) ([]string, error) {
	glob := filepath.Join(dire, "*"+extn)
	origs, err := filepath.Glob(glob)
	if err != nil {
		return nil, fmt.Errorf("cannot list directory %q", dire)
	}

	slices.Sort(origs)
	return origs, nil
}

// Join returns a joined path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a path's base name contains a substring.
func Match(orig, term string) bool {
	term = strings.ToLower(term)
	name := strings.ToLower(Name(orig))
	return strings.Contains(name, term)
}

// Name returns a path's base name without the extension.
func Name(orig string) string {
	base := filepath.Base(orig)
	return strings.SplitN(base, ".", 2)[0]
}
