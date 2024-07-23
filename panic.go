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

// PanicOf expects the function fn to panic by the expected error. If the function does not panic
// or panic for another reason, it will set the result to fail.
//
//	assertion.PanicOf(func() {
//	  panic("expected error")
//	}, "expected error") // success
//	assertion.PanicOf(func() {
//	  panic("unexpected error")
//	}, "expected error") // fail
//	assertion.PanicOf(func() {
//	  // ..., no panic
//	}, "expected error") // fail
func (a *Assertion) PanicOf(fn func(), expectErr any, message ...any) error {
	a.Helper()

	return tryPanicOf(a.T, false, fn, expectErr, message...)
}

// PanicOfNow expects the function fn to panic by the expected error. If the function does not
// panic or panic for another reason, it will set the result to fail and terminate the execution.
//
//	assertion.PanicOfNow(func() {
//	  panic("expected error")
//	}, "expected error") // success
//	assertion.PanicOfNow(func() {
//	  panic("unexpected error")
//	}, "expected error") // fail and terminated
//	// never runs
func (a *Assertion) PanicOfNow(fn func(), expectErr any, message ...any) error {
	a.Helper()

	return tryPanicOf(a.T, true, fn, expectErr, message...)
}

// NotPanicOf expects the function fn not panic, or the function does not panic by the unexpected
// error. If the function panics by the unexpected error, it will set the result to fail.
//
//	assertion.NotPanicOf(func() {
//	  panic("other error")
//	}, "unexpected error") // success
//	assertion.NotPanicOf(func() {
//	  // ..., no panic
//	}, "unexpected error") // success
//	assertion.NotPanicOf(func() {
//	  panic("unexpected error")
//	}, "unexpected error") // fail
func (a *Assertion) NotPanicOf(fn func(), unexpectedErr any, message ...any) error {
	a.Helper()

	return tryNotPanicOf(a.T, false, fn, unexpectedErr, message...)
}

// NotPanicOfNow expects the function fn not panic, or the function does not panic by the
// unexpected error. If the function panics by the unexpected error, it will set the result to fail
// and stop the execution.
//
//	assertion.NotPanicOfNow(func() {
//	  panic("other error")
//	}, "unexpected error") // success
//	assertion.NotPanicOfNow(func() {
//	  // ..., no panic
//	}, "unexpected error") // success
//	assertion.NotPanicOfNow(func() {
//	  panic("unexpected error")
//	}, "unexpected error") // fail and terminate
//	// never runs
func (a *Assertion) NotPanicOfNow(fn func(), unexpectedErr any, message ...any) error {
	a.Helper()

	return tryNotPanicOf(a.T, true, fn, unexpectedErr, message...)
}

// tryPanicOf executes the function fn, and it expects the function to panic by the expected error.
func tryPanicOf(t *testing.T, failedNow bool, fn func(), expectError any, message ...any) error {
	t.Helper()

	e := isPanic(fn)
	if isEqual(e, expectError) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf(defaultErrMessagePanicOf, expectError, e), message...)
	failed(t, err, failedNow)

	return err
}

func tryNotPanicOf(
	t *testing.T,
	failedNow bool,
	fn func(),
	unexpectedError any,
	message ...any,
) error {
	t.Helper()

	e := isPanic(fn)
	if !isEqual(e, unexpectedError) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf(defaultErrMessageNotPanicOf, unexpectedError), message...)
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
