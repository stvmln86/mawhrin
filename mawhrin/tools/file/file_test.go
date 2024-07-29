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
	err := Create(dest, "Body.\n", 0666)
	test.AssertFile(t, dest, "Body.\n")
	assert.NoError(t, err)

	// failure - already exists
	err = Create(dest, "Body.\n", 0666)
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
	nope := filepath.Join(t.TempDir(), "create.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - nope
	ok = Exists(nope)
	assert.False(t, ok)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := filepath.Join(t.TempDir(), "create.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)

	// failure - does not exist
	body, err = Read(nope)
	assert.Empty(t, body)
	test.AssertErr(t, err, "cannot read file .* - does not exist")
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := filepath.Join(t.TempDir(), "create.extn")

	// success
	err := Update(orig, "Body.\n", 0666)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)

	// failure - does not exist
	err = Update(nope, "Body.\n", 0666)
	test.AssertErr(t, err, "cannot update file .* - does not exist")
}
