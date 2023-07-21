package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// DeepEqual tests deeply equality between actual and expect parameters.
func (a *Assertion) DeepEqual(actual, expect any, message ...string) error {
	return tryDeepEqual(a.t, false, actual, expect, message...)
}

// NotDeepEqual tests deeply inequality between actual and expected parameters.
func (a *Assertion) NotDeepEqual(actual, expect any, message ...string) error {
	return tryNotDeepEqual(a.t, false, actual, expect, message...)
}

// tryDeepEqual try to testing the deeply equality between actual and expect values, and it'll
// fail if the values are not deeply equal.
func tryDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
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
	return tryNil(a.t, false, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
func (a *Assertion) NotNil(val any, message ...string) error {
	return tryNotNil(a.t, false, val, message...)
}

// tryNil try to testing a value is nil or not, and it'll fail the value is nil.
func tryNil(t *testing.T, failedNow bool, val any, message ...string) error {
	if isNil(val) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("expect nil, got %v", val), message...)
	failed(t, err, failedNow)

	return err
}

// tryNotNil try to testing a value is nil or not, and it'll fail the value is not nil.
func tryNotNil(t *testing.T, failedNow bool, val any, message ...string) error {
	if !isNil(val) {
		return nil
	}

	err := newAssertionError("expect not nil, got nil", message...)
	failed(t, err, failedNow)

	return err
}
