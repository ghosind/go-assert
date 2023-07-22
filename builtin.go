package assert

import (
	"testing"
)

// DeepEqual tests the deep equality between actual and expect parameters. It'll set the result to
// fail if they are not deeply equal, and it doesn't stop the execution.
func DeepEqual(t *testing.T, actual, expect any, message ...string) error {
	return tryDeepEqual(t, false, actual, expect, message...)
}

// DeepEqualNow tests the deep equality between actual and expect parameters, and it'll stop the
// execution if they are not deeply equal.
func DeepEqualNow(t *testing.T, actual, expect any, message ...string) error {
	return tryDeepEqual(t, true, actual, expect, message...)
}

// NotDeepEqual tests the deep inequality between actual and expected parameters. It'll set the
// result to fail if they are deeply equal, but it doesn't stop the execution.
func NotDeepEqual(t *testing.T, actual, expect any, message ...string) error {
	return tryNotDeepEqual(t, false, actual, expect, message...)
}

// NotDeepEqualNow tests the deep inequality between actual and expected parameters, and it'll stop
// the execution if they are deeply equal.
func NotDeepEqualNow(t *testing.T, actual, expect any, message ...string) error {
	return tryNotDeepEqual(t, true, actual, expect, message...)
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
func Nil(t *testing.T, val any, message ...string) error {
	return tryNil(t, false, val, message...)
}

// NilNow tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is not nil.
func NilNow(t *testing.T, val any, message ...string) error {
	return tryNil(t, true, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
func NotNil(t *testing.T, val any, message ...string) error {
	return tryNotNil(t, false, val, message...)
}

// NotNilNow tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is nil.
func NotNilNow(t *testing.T, val any, message ...string) error {
	return tryNotNil(t, true, val, message...)
}

// Panic expects the function fn to panic, and it'll set the result to fail if the function doesn't
// panic.
func Panic(t *testing.T, fn func(), message ...string) error {
	return tryPanic(t, false, fn, message...)
}

// PanicNow expects the function fn to panic. It'll set the result to fail if the function doesn't
// panic, and stop the execution.
func PanicNow(t *testing.T, fn func(), message ...string) error {
	return tryPanic(t, true, fn, message...)
}

// NotPanic asserts that the function fn does not panic, and it'll set the result to fail if the
// function panic.
func NotPanic(t *testing.T, fn func(), message ...string) error {
	return tryNotPanic(t, false, fn, message...)
}

// NotPanicNow asserts that the function fn does not panic. It'll set the result to fail if the
// function panic, and it also stops the execution.
func NotPanicNow(t *testing.T, fn func(), message ...string) error {
	return tryNotPanic(t, false, fn, message...)
}
