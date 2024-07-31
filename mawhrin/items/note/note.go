// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/stvmln86/mawhrin/mawhrin/tools/file"
	"github.com/stvmln86/mawhrin/mawhrin/tools/neat"
	"github.com/stvmln86/mawhrin/mawhrin/tools/path"
)

// Note is a single plaintext file in a directory.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{orig, mode}
}

// Delete moves the Note to the ".trash" extension.
func (n *Note) Delete() error {
	return file.Delete(n.Orig)
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Match returns true if the Note's name contains a substring.
func (n *Note) Match(term string) bool {
	return path.Match(n.Orig, term)
}

// Name returns the Note's name.
func (n *Note) Name() string {
	return neat.Name(path.Name(n.Orig))
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	body, err := file.Read(n.Orig)
	return neat.Body(body), err
}

// Search returns true if the Note's body contains a string.
func (n *Note) Search(term string) (bool, error) {
	return file.Search(n.Orig, term)
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	return file.Update(n.Orig, body, n.Mode)
}
