// Package book implements the Book type and methods.
package book

import (
	"os"

	"github.com/stvmln86/mawhrin/mawhrin/items/note"
	"github.com/stvmln86/mawhrin/mawhrin/tools/file"
	"github.com/stvmln86/mawhrin/mawhrin/tools/neat"
	"github.com/stvmln86/mawhrin/mawhrin/tools/path"
)

// Book is a single directory containing multiple Notes.
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode os.FileMode) *Book {
	dire = neat.Path(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// Create returns a new Note containing a string.
func (b *Book) Create(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	body = neat.Body(body)
	dest := path.Join(b.Dire, name, b.Extn)
	if err := file.Create(dest, body, b.Mode); err != nil {
		return nil, err
	}

	return note.New(dest, b.Mode), nil
}

// CreateOrGet returns a new (and empty) or existing Note.
func (b *Book) CreateOrGet(name string) (*note.Note, error) {
	name = neat.Name(name)
	dest := path.Join(b.Dire, name, b.Extn)
	if !file.Exists(dest) {
		if err := file.Create(dest, "", b.Mode); err != nil {
			return nil, err
		}
	}

	return b.Get(name)
}

// Filter returns all existing Notes passing a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	notes, err := b.List()
	if err != nil {
		return nil, err
	}

	var goods []*note.Note
	for _, note := range notes {
		ok, err := ffun(note)
		switch {
		case err != nil:
			return nil, err
		case ok:
			goods = append(goods, note)
		}
	}

	return goods, nil
}

// Get returns an existing Note.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	orig := path.Join(b.Dire, name, b.Extn)
	orig, err := file.Get(orig)
	if err != nil {
		return nil, err
	}

	return note.New(orig, b.Mode), nil
}

// List returns a sorted slice of all existing Notes.
func (b *Book) List() ([]*note.Note, error) {
	origs, err := path.Glob(b.Dire, b.Extn)
	if err != nil {
		return nil, err
	}

	var notes []*note.Note
	for _, orig := range origs {
		notes = append(notes, note.New(orig, b.Mode))
	}

	return notes, nil
}

// Match returns all existing Notes with names containing a substring.
func (b *Book) Match(term string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(term), nil
	})
}

// Search returns all existing Notes with bodies containing a substring.
func (b *Book) Search(term string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(term)
	})
}
