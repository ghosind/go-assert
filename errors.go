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
func (a *Assertion) IsError(err, expected error, message ...any) error {
	return isError(a.T, false, err, expected, message...)
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
func (a *Assertion) IsErrorNow(err, expected error, message ...any) error {
	return isError(a.T, true, err, expected, message...)
}

// NotIsError tests whether the error matches the target or not. It'll set the result to fail if
// the error matches to the target error, and it doesn't stop the execution.
//
//	a := assert.New(t)
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	err3 := errors.New("error 3")
//	a.NotIsError(err1, err2) // success
//	a.NotIsError(err1, err1) // fail
//	a.NotIsError(errors.Join(err1, err2), err3) // success
//	a.NotIsError(errors.Join(err1, err2), err1) // fail
//	a.NotIsError(errors.Join(err1, err2), err2) // fail
func (a *Assertion) NotIsError(err, unexpected error, message ...any) error {
	return notIsError(a.T, false, err, unexpected, message...)
}

// NotIsErrorNow tests whether the error matches the target or not. It'll set the result to fail
// and stop the execution if the error matches to the target error.
//
//	a := assert.New(t)
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	err3 := errors.new("error 3")
//	a.NotIsErrorNow(errors.Join(err1, err2), err3) // success
//	a.NotIsErrorNow(err1, err2) // fail
//	a.NotIsErrorNow(err1, err1) // fail and terminate
//	// never runs
func (a *Assertion) NotIsErrorNow(err, unexpected error, message ...any) error {
	return notIsError(a.T, true, err, unexpected, message...)
}

// isError tests whether the error matches the target or not.
func isError(t *testing.T, failedNow bool, err, expected error, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return errors.Is(err, expected) },
		failedNow,
		fmt.Sprintf(defaultErrMessageIsError, expected, err),
		message...,
	)
}

// isError tests whether the error does not match the target error or not.
func notIsError(t *testing.T, failedNow bool, err, unexpected error, message ...any) error {
	return test(
		t,
		func() bool { return !errors.Is(err, unexpected) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotIsError, unexpected),
		message...,
	)
}
