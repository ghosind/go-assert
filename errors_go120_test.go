//go:build go1.20

package assert

import (
	"errors"
	"testing"
)

func TestIsErrorWithJoinedErrors(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	testIsError(a, mockA, errors.Join(err1, err2), err1, true)
	testIsError(a, mockA, errors.Join(err1, err2), err2, true)
	testIsError(a, mockA, errors.Join(err1, err2), err3, false)
}
