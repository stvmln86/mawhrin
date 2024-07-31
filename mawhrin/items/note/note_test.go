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

func TestDelete(t *testing.T) {
	// setup
	note := xNote(t)
	dest := strings.Replace(note.Orig, ".extn", ".trash", 1)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	ok := note.Exists()
	assert.True(t, ok)
}

func TestMatch(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	ok := note.Match("ALPHA")
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	ok, err := note.Search("ALPHA")
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	err := note.Update("Update.\n")
	test.AssertFile(t, note.Orig, "Update.\n")
	assert.NoError(t, err)
}
