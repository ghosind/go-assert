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
