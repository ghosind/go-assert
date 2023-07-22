package assert

import (
	"fmt"
	"testing"
)

// Panic expects the function fn to panic, and it'll set the result to fail if the function doesn't
// panic.
func (a *Assertion) Panic(fn func(), message ...string) error {
	return tryPanic(a.t, false, fn, message...)
}

// PanicNow expects the function fn to panic. It'll set the result to fail if the function doesn't
// panic, and stop the execution.
func (a *Assertion) PanicNow(fn func(), message ...string) error {
	return tryPanic(a.t, true, fn, message...)
}

// NotPanic asserts that the function fn does not panic, and it'll set the result to fail if the
// function panic.
func (a *Assertion) NotPanic(fn func(), message ...string) error {
	return tryNotPanic(a.t, false, fn, message...)
}

// NotPanicNow asserts that the function fn does not panic. It'll set the result to fail if the
// function panic, and it also stops the execution.
func (a *Assertion) NotPanicNow(fn func(), message ...string) error {
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
