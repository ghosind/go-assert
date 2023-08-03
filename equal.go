package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// DeepEqual tests the deep equality between actual and expect parameters. It'll set the result to
// fail if they are not deeply equal, and it doesn't stop the execution.
func (a *Assertion) DeepEqual(actual, expect any, message ...string) error {
	a.Helper()

	return tryDeepEqual(a.T, false, actual, expect, message...)
}

// DeepEqualNow tests the deep equality between actual and expect parameters, and it'll stop the
// execution if they are not deeply equal.
func (a *Assertion) DeepEqualNow(actual, expect any, message ...string) error {
	a.Helper()

	return tryDeepEqual(a.T, true, actual, expect, message...)
}

// NotDeepEqual tests the deep inequality between actual and expected parameters. It'll set the
// result to fail if they are deeply equal, but it doesn't stop the execution.
func (a *Assertion) NotDeepEqual(actual, expect any, message ...string) error {
	a.Helper()

	return tryNotDeepEqual(a.T, false, actual, expect, message...)
}

// NotDeepEqualNow tests the deep inequality between actual and expected parameters, and it'll stop
// the execution if they are deeply equal.
func (a *Assertion) NotDeepEqualNow(actual, expect any, message ...string) error {
	a.Helper()

	return tryNotDeepEqual(a.T, true, actual, expect, message...)
}

// tryDeepEqual try to testing the deeply equality between actual and expect values, and it'll
// fail if the values are not deeply equal.
func tryDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
	t.Helper()

	if reflect.DeepEqual(actual, expect) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("%v == %v", actual, expect), message...)
	failed(t, err, failedNow)

	return err
}

// tryNotDeepEqual try to testing the deeply inequality between actual and expect values, and it'll
// fail if the values are deeply equal.
func tryNotDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
	t.Helper()

	if !reflect.DeepEqual(actual, expect) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("%v != %v", actual, expect), message...)
	failed(t, err, failedNow)

	return err
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
func (a *Assertion) Nil(val any, message ...string) error {
	a.Helper()

	return tryNil(a.T, false, val, message...)
}

// NilNow tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is not nil.
func (a *Assertion) NilNow(val any, message ...string) error {
	a.Helper()

	return tryNil(a.T, true, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
func (a *Assertion) NotNil(val any, message ...string) error {
	a.Helper()

	return tryNotNil(a.T, false, val, message...)
}

// NotNilNow tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is nil.
func (a *Assertion) NotNilNow(val any, message ...string) error {
	a.Helper()

	return tryNotNil(a.T, true, val, message...)
}

// tryNil try to testing a value is nil or not, and it'll fail the value is nil.
func tryNil(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	if isNil(val) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("expect nil, got %v", val), message...)
	failed(t, err, failedNow)

	return err
}

// tryNotNil try to testing a value is nil or not, and it'll fail the value is not nil.
func tryNotNil(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	if !isNil(val) {
		return nil
	}

	err := newAssertionError("expect not nil, got nil", message...)
	failed(t, err, failedNow)

	return err
}

// True tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
func (a *Assertion) True(val any, message ...string) error {
	a.Helper()

	return tryTrue(a.T, false, val, message...)
}

// TrueNow tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
// The function will stop the execution if the value is falsy.
func (a *Assertion) TrueNow(val any, message ...string) error {
	a.Helper()

	return tryTrue(a.T, true, val, message...)
}

// NotTrue tests whether a value is truthy or not. It'll set the result to fail if the value is a
// truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
func (a *Assertion) NotTrue(val any, message ...string) error {
	a.Helper()

	return tryNotTrue(a.T, false, val, message...)
}

// NotTrueNow tests whether a value is truthy or not. It'll set the result to fail if the value is
// a truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
// The function will stop the execution if the value is truthy.
func (a *Assertion) NotTrueNow(val any, message ...string) error {
	a.Helper()

	return tryNotTrue(a.T, true, val, message...)
}

// tryTrue try to testing a value is truthy or falsy, and it'll fail the value is falsy.
func tryTrue(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	if isTrue(val) {
		return nil
	}

	err := newAssertionError("the expression evaluated to a falsy value")
	failed(t, err, failedNow)

	return err
}

// tryNotTrue try to testing a value is truthy or falsy, and it'll fail the value is truthy.
func tryNotTrue(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	if !isTrue(val) {
		return nil
	}

	err := newAssertionError("the expression evaluated to a truthy value")
	failed(t, err, failedNow)

	return err
}
