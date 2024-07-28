package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestDate(t *testing.T) {
	// setup
	tobj := time.Date(2000, time.January, 2, 0, 0, 0, 0, time.UTC)

	// success
	date := Date(tobj)
	assert.Equal(t, "2000-01-02", date)
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
	path := Path("/././path")
	assert.Equal(t, "/path", path)
}

func TestTime(t *testing.T) {
	// setup
	expe := time.Date(2000, time.January, 2, 0, 0, 0, 0, time.UTC)

	// success
	tobj := Time("2000-01-02")
	assert.Equal(t, expe, tobj)
}
