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
