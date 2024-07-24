package assert

import (
	"errors"
	"fmt"
	"testing"
)

// IsError tests whether the error matches the target or not. It'll set the result to fail if the
// error does not match to the target error, and it doesn't stop the execution.
//
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	a := assert.New(t)
//	a.IsError(err1, err1) // success
//	a.IsError(err1, err2) // fail
//	a.IsError(errors.Join(err1, err2), err1) // success
//	a.IsError(errors.Join(err1, err2), err2) // success
func (a *Assertion) IsError(err, target error, message ...any) error {
	return isError(a.T, false, err, target, message...)
}

// IsErrorNow tests whether the error matches the target or not. It'll set the result to fail and
// stop the execution if the error does not match to the target error.
//
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	a := assert.New(t)
//	a.IsErrorNow(errors.Join(err1, err2), err1) // success
//	a.IsErrorNow(errors.Join(err1, err2), err2) // success
//	a.IsErrorNow(err1, err1) // success
//	a.IsErrorNow(err1, err2) // fail
//	// never runs
func (a *Assertion) IsErrorNow(err, target error, message ...any) error {
	return isError(a.T, true, err, target, message...)
}

// isError tests whether the error matches the target or not.
func isError(t *testing.T, failedNow bool, err, target error, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return errors.Is(err, target) },
		failedNow,
		fmt.Sprintf(defaultErrMessageIsError, target, err),
		message...,
	)
}
