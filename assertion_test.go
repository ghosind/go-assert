package assert

import (
	"testing"
)

func TestNewAssertion(t *testing.T) {
	Panic(t, func() {
		New(nil)
	})

	NotPanic(t, func() {
		New(new(testing.T))
	})
}

func TestRun(t *testing.T) {
	a := New(t)
	isSubTestRun := false

	a.Run("sub test", func(sub *Assertion) {
		DeepEqualNow(t, sub.Name(), "TestRun/sub_test")
		isSubTestRun = true
	})

	DeepEqualNow(t, isSubTestRun, true)
}
