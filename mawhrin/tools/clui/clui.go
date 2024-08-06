// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// GetEnv returns an existing environment variable.
func GetEnv(name string) (string, error) {
	evar, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("environment variable %q does not exist", name)
	}

	return strings.TrimSpace(evar), nil
}

// Parse returns an argument map from parameter and argument slices. If a parameter
// contains a colon, the text after the colon is a default value.
func Parse(paras, cargs []string) (map[string]string, error) {
	var amap = make(map[string]string, len(paras))
	for n, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")
		switch {
		case n >= len(cargs) && ok:
			amap[name] = dflt
		case n >= len(cargs) && !ok:
			return nil, fmt.Errorf("argument %q not specified", name)
		default:
			amap[name] = cargs[n]
		}
	}

	return amap, nil
}

// Split returns a command name and argument slice from an argument slice.
func Split(cargs []string) (string, []string) {
	switch len(cargs) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(cargs[0]), nil
	default:
		return strings.ToLower(cargs[0]), cargs[1:]
	}
}
