package assert

import (
	"fmt"
	"testing"
)

// Panic expects the function fn to panic.
func (a *Assertion) Panic(fn func(), message ...string) error {
	return tryPanic(a.t, false, fn, message...)
}

// NotPanic asserts that the function fn does not panic.
func (a *Assertion) NotPanic(fn func(), message ...string) error {
	return tryNotPanic(a.t, false, fn, message...)
}

// tryPanic executes the function fn, and try to catching the panic error. It expect the function
// fn to panic, and returns error if fn does not panic.
func tryPanic(t *testing.T, failedNow bool, fn func(), message ...string) error {
	e := isPanic(fn)
	if e != nil {
		return nil
	}

	err := newAssertionError("missing expected panic", message...)
	failed(t, err, failedNow)

	return err
}

// tryNotPanic executes the function fn, and try to catching the panic error. It expect the
// function fn does not to panic, and returns error if panic.
func tryNotPanic(t *testing.T, failedNow bool, fn func(), message ...string) error {
	e := isPanic(fn)
	if e == nil {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("got unwanted error: %v", e), message...)
	failed(t, err, failedNow)
	return err
}
