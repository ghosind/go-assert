package assert

import (
	"errors"
	"testing"
)

func TestIsError(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	testIsError(a, mockA, errors.New("test"), err1, false)
	testIsError(a, mockA, err1, err1, true)
	testIsError(a, mockA, err1, err2, false)
	testIsError(a, mockA, err1, err3, false)
}

func testIsError(a, mockA *Assertion, err, target error, isError bool) {
	a.T.Helper()

	testAssertionFunction(a, "IsError", func() error {
		return IsError(mockA.T, err, target)
	}, isError)
	testAssertionFunction(a, "Assertion.IsError", func() error {
		return mockA.IsError(err, target)
	}, isError)

	testAssertionFunction(a, "NotIsError", func() error {
		return NotIsError(mockA.T, err, target)
	}, !isError)
	testAssertionFunction(a, "Assertion.NotIsError", func() error {
		return mockA.NotIsError(err, target)
	}, !isError)

	testAssertionNowFunction(a, "IsErrorNow", func() {
		IsErrorNow(mockA.T, err, target)
	}, !isError)
	testAssertionNowFunction(a, "Assertion.IsErrorNow", func() {
		mockA.IsErrorNow(err, target)
	}, !isError)

	testAssertionNowFunction(a, "NotIsErrorNow", func() {
		NotIsErrorNow(mockA.T, err, target)
	}, isError)
	testAssertionNowFunction(a, "Assertion.NotIsErrorNow", func() {
		mockA.NotIsErrorNow(err, target)
	}, isError)
}
