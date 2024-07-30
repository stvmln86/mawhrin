// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
)

// Body returns a whitespace-trimmed file body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Extn returns a lowercase file extension string with a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a lowercase file name string.
func Name(name string) string {
	return strings.ToLower(name)
}

// Path returns a clean file path.
func Path(orig string) string {
	return filepath.Clean(orig)
}
