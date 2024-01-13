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
func (a *Assertion) Panic(fn func(), message ...any) error {
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
func (a *Assertion) PanicNow(fn func(), message ...any) error {
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
func (a *Assertion) NotPanic(fn func(), message ...any) error {
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
func (a *Assertion) NotPanicNow(fn func(), message ...any) error {
	a.Helper()

	return tryNotPanic(a.T, true, fn, message...)
}

// tryPanic executes the function fn, and try to catching the panic error. It expect the function
// fn to panic, and returns error if fn does not panic.
func tryPanic(t *testing.T, failedNow bool, fn func(), message ...any) error {
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
func tryNotPanic(t *testing.T, failedNow bool, fn func(), message ...any) error {
	t.Helper()

	e := isPanic(fn)
	if e == nil {
		return nil
	}

	err := newAssertionError(fmt.Sprintf(defaultErrMessageNotPanic, e), message...)
	failed(t, err, failedNow)
	return err
}

// PanicOf expects the function fn to panic by the expected error.
func (a *Assertion) PanicOf(fn func(), expectErr any, message ...any) error {
	a.Helper()

	return tryPanicOf(a.T, false, fn, expectErr, message...)
}

// tryPanicOf executes the function fn, and it expects the function to panic by the expected error.
func tryPanicOf(t *testing.T, failedNow bool, fn func(), expectError any, message ...any) error {
	t.Helper()

	e := isPanic(fn)
	if isEqual(e, expectError) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf(defaultErrMessagePanicOf, expectError, e))
	failed(t, err, failedNow)

	return err
}

// isPanic executes the function, and tries to catching and returns the return value from
// recover().
func isPanic(fn func()) (err any) {
	defer func() {
		if e := recover(); e != nil {
			err = e
		}
	}()

	fn()

	return
}
