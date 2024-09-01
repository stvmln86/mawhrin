// Package tasks implements the Task interface and types.
package tasks

import (
	"fmt"
	"io"
	"strings"

	"github.com/stvmln86/mawhrin/mawhrin/items/book"
	"github.com/stvmln86/mawhrin/mawhrin/tools/clui"
)

// Task is a user-facing command-line function.
type Task struct {
	Name  string
	Paras []string
	Help  string
	Func  TaskFunc
}

// TaskFunc is the underlying function of a Task.
type TaskFunc func(io.Writer, *book.Book, map[string]string) error

// New returns a new initialised Task.
func New(lede, help string, tfun TaskFunc) *Task {
	name, paras, _ := strings.Cut(lede, " ")
	return &Task{name, strings.Fields(paras), help, tfun}
}

// Tasks is a slice of existing Tasks.
var Tasks = map[string]*Task{
	"list": ListTask,
}

// Run executes an existing Task from an argument slice.
func Run(w io.Writer, book *book.Book, cargs []string) error {
	name, cargs := clui.Split(cargs)
	task, ok := Tasks[name]
	if !ok {
		return fmt.Errorf("task %q does not exist", name)
	}

	amap, err := clui.Parse(task.Paras, cargs)
	if err != nil {
		return err
	}

	return task.Func(w, book, amap)
}
