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

// Equal tests equality between actual and expect parameters.
func (a *Assertion) Equal(actual, expect any, message ...string) error {
	return Equal(a.t, actual, expect, message...)
}

// NotEqual tests inequality between actual and expected parameters.
func (a *Assertion) NotEqual(actual, expect any, message ...string) error {
	return NotEqual(a.t, actual, expect, message...)
}

// Panic expects the function fn to panic.
func (a *Assertion) Panic(fn func(), message ...string) (err error) {
	return Panic(a.t, fn, message...)
}

// NotPanic asserts that the function fn does not panic.
func (a *Assertion) NotPanic(fn func(), message ...string) (err error) {
	return NotPanic(a.t, fn, message...)
}
