// Package tasks implements the Task interface and types.
package tasks

import (
	"fmt"
	"io"

	"github.com/stvmln86/mawhrin/mawhrin/items/book"
	"github.com/stvmln86/mawhrin/mawhrin/tools/clui"
)

// Task is a user-facing command-line function.
type Task interface {
	// Name returns the Task's name.
	Name() string

	// Help returns the Task's help string.
	Help() string

	// Paras returns the Task's argument parameters.
	Paras() []string

	// Run executes the Task.
	Run(io.Writer, map[string]string) error
}

// NewTask returns a new initialised Task.
type NewTask func(*book.Book) (Task, error)

// Tasks is a slice of existing Tasks.
var Tasks = map[string]NewTask{
	"list": NewListTask,
}

// Run executes an existing Task from an argument slice.
func Run(w io.Writer, book *book.Book, cargs []string) error {
	name, cargs := clui.Split(cargs)
	tfun, ok := Tasks[name]
	if !ok {
		return fmt.Errorf("command %q does not exist", name)
	}

	task, err := tfun(book)
	if err != nil {
		return err
	}

	amap, err := clui.Parse(task.Paras(), cargs)
	if err != nil {
		return err
	}

	return task.Run(w, amap)
}
