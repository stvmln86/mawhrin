package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	// success
	body := join("alpha", "bravo")
	assert.Equal(t, "alpha\nbravo\n", body)
}

func TestAssertDire(t *testing.T) {
	// setup
	dire := MockDire(t)

	// success
	AssertDire(t, dire, MockData)
}

func TestAssertErr(t *testing.T) {
	// setup
	err := errors.New("error")

	// success
	AssertErr(t, err, ".*")
}

func TestAssertFile(t *testing.T) {
	// setup
	orig := MockFile(t, "alpha.extn")

	// success
	AssertFile(t, orig, MockData["alpha.extn"])
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	AssertDire(t, dire, MockData)
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	AssertFile(t, orig, MockData["alpha.extn"])
}
