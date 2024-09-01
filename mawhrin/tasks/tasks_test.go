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

var MockTask = New(
	"mock PARAMETER",
	"Mock task.",

	func(w io.Writer, book *book.Book, amap map[string]string) error {
		fmt.Fprintf(w, "PARAMETER=%s", amap["PARAMETER"])
		return nil
	},
)

func runTask(t *testing.T, cargs string) (*bytes.Buffer, error) {
	w := bytes.NewBuffer(nil)
	dire := test.MockDire(t)
	book := book.New(dire, ".extn", 0666)
	return w, Run(w, book, strings.Split(cargs, " "))
}

func TestNew(t *testing.T) {
	// success
	assert.Equal(t, "mock", MockTask.Name)
	assert.Equal(t, []string{"PARAMETER"}, MockTask.Paras)
	assert.Equal(t, "Mock task.", MockTask.Help)
	assert.NotNil(t, MockTask.Func)
}

func TestRun(t *testing.T) {
	// setup
	Tasks["mock"] = MockTask

	// success
	w, err := runTask(t, "mock argument")
	assert.Equal(t, "PARAMETER=argument", w.String())
	assert.NoError(t, err)

	// failure
	w, err = runTask(t, "nope argument")
	assert.Empty(t, w.String())
	test.AssertErr(t, err, `task "nope" does not exist`)
}
