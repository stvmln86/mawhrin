package tasks

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/mawhrin/mawhrin/items/book"
	"github.com/stvmln86/mawhrin/mawhrin/tools/test"
)

type MockTask struct{ Book *book.Book }

func NewMockTask(book *book.Book) (Task, error) {
	return &MockTask{book}, nil
}

func (t *MockTask) Name() string {
	return "mock"
}

func (t *MockTask) Help() string {
	return "Mock task."
}

func (t *MockTask) Paras() []string {
	return []string{"PARAMETER"}
}

func (t *MockTask) Run(w io.Writer, amap map[string]string) error {
	fmt.Fprintf(w, "PARAMETER=%s", amap["PARAMETER"])
	return nil
}

func runTask(t *testing.T, cargs string) (*bytes.Buffer, error) {
	w := bytes.NewBuffer(nil)
	dire := test.MockDire(t)
	book := book.New(dire, ".extn", 0666)
	return w, Run(w, book, strings.Split(cargs, " "))
}

func TestRun(t *testing.T) {
	// setup
	Tasks["mock"] = NewMockTask

	// success
	w, err := runTask(t, "mock argument")
	assert.Equal(t, "PARAMETER=argument", w.String())
	assert.NoError(t, err)

	// failure
	w, err = runTask(t, "nope argument")
	assert.Empty(t, w.String())
	test.AssertErr(t, err, `command "nope" does not exist`)
}
