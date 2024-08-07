package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTask(t *testing.T) {
	// success - no argument
	w, err := runTask(t, "list")
	assert.Equal(t, "alpha\nbravo\n", w.String())
	assert.NoError(t, err)

	// success - with argument
	w, err = runTask(t, "list alpha")
	assert.Equal(t, "alpha\n", w.String())
	assert.NoError(t, err)
}
