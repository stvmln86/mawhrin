package tasks

import (
	"fmt"
	"io"

	"github.com/stvmln86/mawhrin/mawhrin/items/book"
)

// ListTask is a Task that lists Notes.
var ListTask = New(
	"list TEXT:",
	"List all notes, or notes matching TEXT.",

	func(w io.Writer, book *book.Book, amap map[string]string) error {
		notes, err := book.Match(amap["TEXT"])
		if err != nil {
			return err
		}

		for _, note := range notes {
			fmt.Fprintf(w, "%s\n", note.Name())
		}

		return nil
	},
)
