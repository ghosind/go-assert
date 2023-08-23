package assert

import (
	"fmt"
	"testing"
)

// Panic expects the function fn to panic, and it'll set the result to fail if the function doesn't
// panic.
//
//	assertion.Panic(func() {
//	  panic("some error")
//	}) // success
//
//	assertion.Panic(func() {
//	  // no panic
//	}) // fail
func (a *Assertion) Panic(fn func(), message ...string) error {
	a.Helper()

	return tryPanic(a.T, false, fn, message...)
}

// PanicNow expects the function fn to panic. It'll set the result to fail if the function doesn't
// panic, and stop the execution.
//
//	assertion.PanicNow(func() {
//	  panic("some error")
//	}) // success
//
//	assertion.PanicNow(func() {
//	  // no panic
//	}) // fail
//	// never run
func (a *Assertion) PanicNow(fn func(), message ...string) error {
	a.Helper()

	return tryPanic(a.T, true, fn, message...)
}

// NotPanic asserts that the function fn does not panic, and it'll set the result to fail if the
// function panic.
//
//	assertion.NotPanic(func() {
//	  // no panic
//	}) // success
//
//	assertion.NotPanic(func() {
//	  panic("some error")
//	}) // fail
func (a *Assertion) NotPanic(fn func(), message ...string) error {
	a.Helper()

	return tryNotPanic(a.T, false, fn, message...)
}

// NotPanicNow asserts that the function fn does not panic. It'll set the result to fail if the
// function panic, and it also stops the execution.
//
//	assertion.NotPanicNow(func() {
//	  // no panic
//	}) // success
//
//	assertion.NotPanicNow(func() {
//	  panic("some error")
//	}) // fail and terminate
//	// never run
func (a *Assertion) NotPanicNow(fn func(), message ...string) error {
	a.Helper()

	return tryNotPanic(a.T, true, fn, message...)
}

// tryPanic executes the function fn, and try to catching the panic error. It expect the function
// fn to panic, and returns error if fn does not panic.
func tryPanic(t *testing.T, failedNow bool, fn func(), message ...string) error {
	t.Helper()

	e := isPanic(fn)
	if e != nil {
		return nil
	}

	err := newAssertionError(defaultErrMessagePanic, message...)
	failed(t, err, failedNow)

	return err
}

// tryNotPanic executes the function fn, and try to catching the panic error. It expect the
// function fn does not to panic, and returns error if panic.
func tryNotPanic(t *testing.T, failedNow bool, fn func(), message ...string) error {
	t.Helper()

	e := isPanic(fn)
	if e == nil {
		return nil
	}

	err := newAssertionError(fmt.Sprintf(defaultErrMessageNotPanic, e), message...)
	failed(t, err, failedNow)
	return err
}
