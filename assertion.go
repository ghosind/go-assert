package assert

import (
	"testing"
)

type Assertion struct {
	*testing.T
}

// New returns an assertion instance for verifying invariants.
func New(t *testing.T) *Assertion {
	a := new(Assertion)

	if t == nil {
		panic("parameter t is required")
	}
	a.T = t

	return a
}

// Run runs f as a subtest of a called name. It runs f in a separate goroutine
// and blocks until f returns or calls a.Parallel to become a parallel test.
// Run reports whether f succeeded (or at least did not fail before calling t.Parallel).
//
// Run may be called simultaneously from multiple goroutines, but all such calls
// must return before the outer test function for a returns.
func (assertion *Assertion) Run(name string, f func(a *Assertion)) bool {
	return assertion.T.Run(name, func(t *testing.T) {
		subAssertion := New(t)
		f(subAssertion)
	})
}
