package book

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/mawhrin/mawhrin/items/note"
	"github.com/stvmln86/mawhrin/mawhrin/tools/test"
)

func xBook(t *testing.T) *Book {
	dire := test.MockDire(t)
	return New(dire, ".extn", 0666)
}

func TestNew(t *testing.T) {
	// success
	book := xBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0666), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := xBook(t)
	dest := filepath.Join(book.Dire, "create.extn")

	// success
	note, err := book.Create("create", "Create.\n")
	assert.Equal(t, dest, note.Orig)
	test.AssertFile(t, note.Orig, "Create.\n")
	assert.NoError(t, err)
}

func TestCreateOrGet(t *testing.T) {
	// setup
	book := xBook(t)
	orig := filepath.Join(book.Dire, "alpha.extn")
	dest := filepath.Join(book.Dire, "create.extn")

	// success - existing note
	note, err := book.CreateOrGet("alpha")
	assert.Equal(t, orig, note.Orig)
	assert.NoError(t, err)

	// success - created note
	note, err = book.CreateOrGet("create")
	assert.Equal(t, dest, note.Orig)
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := xBook(t)
	orig := filepath.Join(book.Dire, "alpha.extn")

	// success
	notes, err := book.Filter(func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	})
	assert.Len(t, notes, 1)
	assert.Equal(t, orig, notes[0].Orig)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := xBook(t)
	orig := filepath.Join(book.Dire, "alpha.extn")

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, orig, note.Orig)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.List()
	assert.Len(t, notes, 2)
	assert.NoError(t, err)
}

func TestMatch(t *testing.T) {
	// setup
	book := xBook(t)
	orig := filepath.Join(book.Dire, "alpha.extn")

	// success
	notes, err := book.Match("ALPHA")
	assert.Len(t, notes, 1)
	assert.Equal(t, orig, notes[0].Orig)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	book := xBook(t)
	orig := filepath.Join(book.Dire, "alpha.extn")

	// success
	notes, err := book.Search("ALPHA")
	assert.Len(t, notes, 1)
	assert.Equal(t, orig, notes[0].Orig)
	assert.NoError(t, err)
}
