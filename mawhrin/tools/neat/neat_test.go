package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("EXTN")
	assert.Equal(t, ".extn", extn)
}

func TestName(t *testing.T) {
	name := Name("NAME")
	assert.Equal(t, "name", name)
}

func TestPath(t *testing.T) {
	// success
	orig := Path("/././path")
	assert.Equal(t, "/path", orig)
}
