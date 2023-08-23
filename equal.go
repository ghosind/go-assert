package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// DeepEqual tests the deep equality between actual and expect parameters. It'll set the result to
// fail if they are not deeply equal, and it doesn't stop the execution.
//
//	assertion.DeepEqual(1, 1) // success
//	assertion.DeepEqual("ABC", "ABC") // success
//	assertion.DeepEqual(1, 0) // fail
//	assertion.DeepEqual(1, int64(1)) // fail
func (a *Assertion) DeepEqual(actual, expect any, message ...string) error {
	a.Helper()

	return tryDeepEqual(a.T, false, actual, expect, message...)
}

// DeepEqualNow tests the deep equality between actual and expect parameters, and it'll stop the
// execution if they are not deeply equal.
//
//	assertion.DeepEqualNow(1, 1) // success
//	assertion.DeepEqualNow("ABC", "ABC") // success
//	assertion.DeepEqualNow(1, int64(1)) // fail and terminate
//	// never run
func (a *Assertion) DeepEqualNow(actual, expect any, message ...string) error {
	a.Helper()

	return tryDeepEqual(a.T, true, actual, expect, message...)
}

// NotDeepEqual tests the deep inequality between actual and expected parameters. It'll set the
// result to fail if they are deeply equal, but it doesn't stop the execution.
//
//	assertion.NotDeepEqual(1, 0) // success
//	assertion.NotDeepEqual(1, int64(1)) // success
//	assertion.NotDeepEqual(1, 1) // fail
//	assertion.NotDeepEqual("ABC", "ABC") // fail
func (a *Assertion) NotDeepEqual(actual, expect any, message ...string) error {
	a.Helper()

	return tryNotDeepEqual(a.T, false, actual, expect, message...)
}

// NotDeepEqualNow tests the deep inequality between actual and expected parameters, and it'll stop
// the execution if they are deeply equal.
//
//	assertion.NotDeepEqual1, 0) // success
//	assertion.NotDeepEqual1, int64(1)) // success
//	assertion.NotDeepEqual"ABC", "ABC") // fail and terminate
//	// never run
func (a *Assertion) NotDeepEqualNow(actual, expect any, message ...string) error {
	a.Helper()

	return tryNotDeepEqual(a.T, true, actual, expect, message...)
}

// tryDeepEqual try to testing the deeply equality between actual and expect values, and it'll
// fail if the values are not deeply equal.
func tryDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool { return reflect.DeepEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageEqual, actual, expect),
		message...,
	)
}

// tryNotDeepEqual try to testing the deeply inequality between actual and expect values, and it'll
// fail if the values are deeply equal.
func tryNotDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool { return !reflect.DeepEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotEqual, actual, expect),
		message...,
	)
}

// Equal tests the equality between actual and expect parameters. It'll set the result to fail if
// they are not equal, and it doesn't stop the execution.
//
//	assertion.Equal(1, 1) // success
//	assertion.Equal("ABC", "ABC") // success
//	assertion.Equal(1, int64(1)) // success
//	assertion.Equal(1, uint64(1)) // success
//	assertion.Equal(1, 0) // fail
func (a *Assertion) Equal(actual, expect any, message ...string) error {
	a.Helper()

	return tryEqual(a.T, false, actual, expect, message...)
}

// EqualNow tests the equality between actual and expect parameters, and it'll stop the execution
// if they are not equal.
//
//	assertion.EqualNow(1, 1) // success
//	assertion.EqualNow("ABC", "ABC") // success
//	assertion.EqualNow(1, int64(1)) // success
//	assertion.EqualNow(1, uint64(1)) // success
//	assertion.EqualNow(1, 0) // fail and terminate
//	never run
func (a *Assertion) EqualNow(actual, expect any, message ...string) error {
	a.Helper()

	return tryEqual(a.T, true, actual, expect, message...)
}

// NotEqual tests the inequality between actual and expected parameters. It'll set the result to
// fail if they are equal, but it doesn't stop the execution.
//
//	assertion.NotEqual(1, 0) // success
//	assertion.NotEqual("ABC", "CBA") // success
//	assertion.NotEqual(1, 1) // fail
//	assertion.NotEqual("ABC", "ABC") // fail
//	assertion.NotEqual(1, int64(1)) // fail
//	assertion.NotEqual(1, uint64(1)) // fail
func (a *Assertion) NotEqual(actual, expect any, message ...string) error {
	a.Helper()

	return tryNotEqual(a.T, false, actual, expect, message...)
}

// NotEqualNow tests the inequality between actual and expected parameters, and it'll stop the
// execution if they are equal.
//
//	assertion.NotEqualNow(1, 0) // success
//	assertion.NotEqualNow("ABC", "CBA") // success
//	assertion.NotEqualNow(1, 1) // fail and terminate
//	// never run
func (a *Assertion) NotEqualNow(actual, expect any, message ...string) error {
	a.Helper()

	return tryNotEqual(a.T, true, actual, expect, message...)
}

// tryEqual try to testing the equality between actual and expect values, and it'll fail if the
// values are not equal.
func tryEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool { return isEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageEqual, actual, expect),
		message...,
	)
}

// tryNotEqual try to testing the inequality between actual and expect values, and it'll fail if
// the values are equal.
func tryNotEqual(t *testing.T, failedNow bool, actual, expect any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotEqual, actual, expect),
		message...,
	)
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
//	var err error // nil
//	assertion.Nil(err) // success
//
//	err = errors.New("some error")
//	assertion.Nil(err) // fail
func (a *Assertion) Nil(val any, message ...string) error {
	a.Helper()

	return tryNil(a.T, false, val, message...)
}

// NilNow tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is not nil.
//
//	var err error // nil
//	assertion.NilNow(err) // success
//
//	err = errors.New("some error")
//	assertion.NilNow(err) // fail and terminate
//	// never run
func (a *Assertion) NilNow(val any, message ...string) error {
	a.Helper()

	return tryNil(a.T, true, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
//	var err error // nil
//	assertion.NotNil(err) // fail
//
//	err = errors.New("some error")
//	assertion.NotNil(err) // success
func (a *Assertion) NotNil(val any, message ...string) error {
	a.Helper()

	return tryNotNil(a.T, false, val, message...)
}

// NotNilNow tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is nil.
//
//	var err error = errors.New("some error")
//	assertion.NotNilNow(err) // success
//
//	err = nil
//	assertion.NotNilNow(err) // fail and terminate
//	// never run
func (a *Assertion) NotNilNow(val any, message ...string) error {
	a.Helper()

	return tryNotNil(a.T, true, val, message...)
}

// tryNil try to testing a value is nil or not, and it'll fail the value is nil.
func tryNil(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool { return isNil(val) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNil, val),
		message...,
	)
}

// tryNotNil try to testing a value is nil or not, and it'll fail the value is not nil.
func tryNotNil(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isNil(val) },
		failedNow,
		defaultErrMessageNotNil,
		message...,
	)
}

// True tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
//	assertion.True(1) // success
//	assertion.True("test") // success
//	assertion.True(0) // fail
//	assertion.True("") // fail
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
//
//	assertion.TrueNow(1) // success
//	assertion.TrueNow("test") // success
//	assertion.TrueNow("") // fail and terminate
//	// never run
func (a *Assertion) TrueNow(val any, message ...string) error {
	a.Helper()

	return tryTrue(a.T, true, val, message...)
}

// NotTrue tests whether a value is truthy or not. It'll set the result to fail if the value is a
// truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
//	assertion.NotTrue(0) // success
//	assertion.NotTrue("") // success
//	assertion.NotTrue(1) // fail
//	assertion.NotTrue("test") // fail
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
//
//	assertion.NotTrueNow(0) // success
//	assertion.NotTrueNow("") // success
//	assertion.NotTrueNow("test") // fail and terminate
//	// never run
func (a *Assertion) NotTrueNow(val any, message ...string) error {
	a.Helper()

	return tryNotTrue(a.T, true, val, message...)
}

// tryTrue try to testing a value is truthy or falsy, and it'll fail the value is falsy.
func tryTrue(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool {
			return isTrue(val)
		},
		failedNow,
		defaultErrMessageTrue,
		message...,
	)
}

// tryNotTrue try to testing a value is truthy or falsy, and it'll fail the value is truthy.
func tryNotTrue(t *testing.T, failedNow bool, val any, message ...string) error {
	t.Helper()

	return test(
		t,
		func() bool {
			return !isTrue(val)
		},
		failedNow,
		defaultErrMessageNotTrue,
		message...,
	)
}
