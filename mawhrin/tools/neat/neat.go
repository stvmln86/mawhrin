// Package neat implements value sanitisation functions.
package neat

import (
	"path/filepath"
	"strings"
	"time"
)

// Body returns a whitespace-trimmed file body string with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Date returns an ISO 8601 date string from a Time object.
func Date(tobj time.Time) string {
	return tobj.Format("2006-01-02")
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
func Path(path string) string {
	return filepath.Clean(path)
}

// Time returns a Time object from an ISO 8601 date string.
func Time(date string) time.Time {
	tobj, _ := time.Parse("2006-01-02", date)
	return tobj
}
