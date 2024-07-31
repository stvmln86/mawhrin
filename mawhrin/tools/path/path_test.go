package path

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/mawhrin/mawhrin/tools/test"
)

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)
}

func TestGlob(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	origs, err := Glob(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, origs)
	assert.NoError(t, err)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", dest)
}

func TestMatch(t *testing.T) {
	// success - true
	ok := Match("/dire/name.extn", "NAME")
	assert.True(t, ok)

	// success - false
	ok = Match("/dire/name.extn", "NOPE")
	assert.False(t, ok)
}

func TestName(t *testing.T) {
	// success
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)
}
