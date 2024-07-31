package note

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/mawhrin/mawhrin/tools/test"
)

func xNote(t *testing.T) *Note {
	orig := test.MockFile(t, "alpha.extn")
	return New(orig, 0666)
}

func TestNew(t *testing.T) {
	// success
	note := xNote(t)
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.Equal(t, os.FileMode(0666), note.Mode)
}

func Delete(t *testing.T) {
	// setup
	note := xNote(t)
	dest := strings.Replace(note.Orig, ".extn", ".trash", 1)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func Match(t *testing.T) {
	// setup
	note := xNote(t)

	// success - true
	ok := note.Match("ALPHA")
	assert.True(t, ok)

	// success - false
	ok = note.Match("NOPE")
	assert.False(t, ok)
}

func Name(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func Read(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)
}

func Search(t *testing.T) {
	// setup
	note := xNote(t)

	// success - true
	ok, err := note.Search("ALPHA")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = note.Search("NOPE")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func Update(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	err := note.Update("Update.\n")
	test.AssertFile(t, note.Orig, "Update.\n")
	assert.NoError(t, err)
}
