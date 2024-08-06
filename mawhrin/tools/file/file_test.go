package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/mawhrin/mawhrin/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dest := filepath.Join(t.TempDir(), "create.extn")

	// success
	err := Create(dest, "Create.\n", 0666)
	test.AssertFile(t, dest, "Create.\n")
	assert.NoError(t, err)

	// failure - already exists
	err = Create(dest, "Create.\n", 0666)
	test.AssertErr(t, err, "cannot create file .* - already exists")
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, ".extn", ".trash", 1)

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Delete(orig)
	test.AssertErr(t, err, "cannot delete file .* - does not exist")
}

func TestExists(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := filepath.Join(t.TempDir(), "nope.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists(nope)
	assert.False(t, ok)
}

func TestGet(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := filepath.Join(t.TempDir(), "nope.extn")

	// success
	dest, err := Get(orig)
	assert.Equal(t, orig, dest)
	assert.NoError(t, err)

	// failure - does not exist
	dest, err = Get(nope)
	assert.Empty(t, dest)
	test.AssertErr(t, err, "cannot get file .* - does not exist")
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := filepath.Join(t.TempDir(), "nope.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)

	// failure - does not exist
	body, err = Read(nope)
	assert.Empty(t, body)
	test.AssertErr(t, err, "cannot read file .* - does not exist")
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha", "rename", 1)

	// success
	err := Rename(orig, "rename")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Rename(orig, "rename")
	test.AssertErr(t, err, "cannot rename file .* - does not exist")
}

func TestSearch(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	ok, err := Search(orig, "ALPHA")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = Search(orig, "NOPE")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := filepath.Join(t.TempDir(), "nope.extn")

	// success
	err := Update(orig, "Update.\n", 0666)
	test.AssertFile(t, orig, "Update.\n")
	assert.NoError(t, err)

	// failure - does not exist
	err = Update(nope, "Update.\n", 0666)
	test.AssertErr(t, err, "cannot update file .* - does not exist")
}
