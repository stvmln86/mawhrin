// Package test implements unit-testing helper functions and data.
package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockData is a base:body map of mock note files.
var MockData = map[string]string{
	"alpha.extn": join(
		".init 2000-01-02",
		".edit 2000-01-03",
		"Alpha.",
	),

	"bravo.extn": join(
		".nope nope",
		"Bravo.",
	),
}

// join returns a string slice as a newline-separated string.
func join(lines ...string) string {
	return strings.Join(lines, "\n") + "\n"
}

// AssertDire asserts the files in a directory are equal to a base:body map.
func AssertDire(t *testing.T, dire string, files map[string]string) {
	for base, body := range files {
		orig := filepath.Join(dire, base)
		AssertFile(t, orig, body)
	}
}

// AssertErr asserts an error's message matches a regular expression.
func AssertErr(t *testing.T, err error, regx string) {
	assert.Regexp(t, regx, err.Error())
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}

// MockDire returns a temporary directory populated with MockData entries.
func MockDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockData {
		dest := filepath.Join(dire, base)
		os.WriteFile(dest, []byte(body), 0666)
	}

	return dire
}

// MockFile returns a temporary file populated with a MockData entry.
func MockFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	os.WriteFile(dest, []byte(MockData[base]), 0666)
	return dest
}
