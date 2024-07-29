package test

import (
	"errors"
	"testing"
)

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
