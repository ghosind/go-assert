package assert

import "testing"

type Assertion struct {
	t *testing.T
}

// New returns an assertion instance for verifying invariants.
func New(t *testing.T) *Assertion {
	a := new(Assertion)

	a.t = t

	return a
}

// DeepEqual tests deeply equality between actual and expect parameters.
func (a *Assertion) DeepEqual(actual, expect any, message ...string) error {
	return DeepEqual(a.t, actual, expect, message...)
}

// NotDeepEqual tests deeply inequality between actual and expected parameters.
func (a *Assertion) NotDeepEqual(actual, expect any, message ...string) error {
	return NotDeepEqual(a.t, actual, expect, message...)
}

// Nil tests a value is nil or not, and it'll failed when the value is not nil.
func (a *Assertion) Nil(val any, message ...string) error {
	return Nil(a.t, val, message...)
}

// NotNil tests a value is nil or not, and it'll failed when the value is nil.
func (a *Assertion) NotNil(val any, message ...string) error {
	return NotNil(a.t, val, message...)
}

// Panic expects the function fn to panic.
func (a *Assertion) Panic(fn func(), message ...string) (err error) {
	return Panic(a.t, fn, message...)
}

// NotPanic asserts that the function fn does not panic.
func (a *Assertion) NotPanic(fn func(), message ...string) (err error) {
	return NotPanic(a.t, fn, message...)
}
