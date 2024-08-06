package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/mawhrin/mawhrin/tools/test"
)

func TestGetEnv(t *testing.T) {
	// setup
	os.Setenv("ENV", "VALUE")

	// success
	evar, err := GetEnv("ENV")
	assert.Equal(t, "VALUE", evar)
	assert.NoError(t, err)

	// failure - does not exist
	evar, err = GetEnv("NOPE")
	assert.Empty(t, evar)
	test.AssertErr(t, err, `environment variable "NOPE" does not exist`)
}

func TestParse(t *testing.T) {
	// success - actual arguments
	amap, err := Parse([]string{"PARAMETER"}, []string{"argument"})
	assert.Equal(t, map[string]string{"PARAMETER": "argument"}, amap)
	assert.NoError(t, err)

	// success - default arguments
	amap, err = Parse([]string{"PARAMETER:default"}, nil)
	assert.Equal(t, map[string]string{"PARAMETER": "default"}, amap)
	assert.NoError(t, err)

	// failure - argument not specified
	amap, err = Parse([]string{"PARAMETER"}, nil)
	assert.Empty(t, amap)
	test.AssertErr(t, err, `argument "PARAMETER" not specified`)
}

func TestSplit(t *testing.T) {
	// success - no arguments
	name, cargs := Split(nil)
	assert.Empty(t, name)
	assert.Empty(t, cargs)

	// success - one arguments
	name, cargs = Split([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Empty(t, cargs)

	// success - multiple arguments
	name, cargs = Split([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, cargs)
}
