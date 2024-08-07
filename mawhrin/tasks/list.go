package tasks

import (
	"fmt"
	"io"

	"github.com/stvmln86/mawhrin/mawhrin/items/book"
)

// ListTask is a Task that lists Notes.
type ListTask struct {
	Book *book.Book
}

// NewListTask returns a new ListTask.
func NewListTask(book *book.Book) (Task, error) {
	return &ListTask{book}, nil
}

// Name returns the Task's name.
func (t *ListTask) Name() string {
	return "list"
}

// Help returns the Task's help string.
func (t *ListTask) Help() string {
	return "List all notes matching SUBSTRING."
}

// Paras returns the Task's argument parameters.
func (t *ListTask) Paras() []string {
	return []string{"SUBSTRING:"}
}

// Run executes the Task.
func (t *ListTask) Run(w io.Writer, amap map[string]string) error {
	notes, err := t.Book.Match(amap["SUBSTRING"])
	if err != nil {
		return err
	}

	for _, note := range notes {
		fmt.Fprintf(w, "%s\n", note.Name())
	}

	return nil
}
