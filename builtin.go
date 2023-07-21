package assert

import (
	"testing"
)

// DeepEqual tests deeply equality between actual and expect parameters.
func DeepEqual(t *testing.T, actual, expect any, message ...string) error {
	return tryDeepEqual(t, false, actual, expect, message...)
}

// NotDeepEqual tests deeply inequality between actual and expected parameters.
func NotDeepEqual(t *testing.T, actual, expect any, message ...string) error {
	return tryNotDeepEqual(t, false, actual, expect, message...)
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
func Nil(t *testing.T, val any, message ...string) error {
	return tryNil(t, false, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
func NotNil(t *testing.T, val any, message ...string) error {
	return tryNotNil(t, false, val, message...)
}

// Panic expects the function fn to panic.
func Panic(t *testing.T, fn func(), message ...string) error {
	return tryPanic(t, false, fn, message...)
}

// NotPanic asserts that the function fn does not panic.
func NotPanic(t *testing.T, fn func(), message ...string) error {
	return tryNotPanic(t, false, fn, message...)
}
