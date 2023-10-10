package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestNewAssertion(t *testing.T) {
	Panic(t, func() {
		New(nil)
	})

	NotPanic(t, func() {
		New(new(testing.T))
	})
}

func TestAssertionWithoutNew(t *testing.T) {
	Panic(t, func() {
		a := new(Assertion)

		a.True(true)
	})
}

func TestRun(t *testing.T) {
	a := New(t)
	isSubTestRun := false

	a.Run("sub test", func(sub *Assertion) {
		EqualNow(t, sub.Name(), "TestRun/sub_test")
		isSubTestRun = true
	})

	TrueNow(t, isSubTestRun)
}

func testAssertionFunction(a *Assertion, name string, fn func() error, expectSuccess bool) {
	a.Helper()

	err := fn()
	if expectSuccess {
		a.NilNow(err, "%s() = %v, want = nil", name, err)
	} else {
		a.NotNilNow(err, "%s() = nil, want = error", name)
	}
}

func testAssertionNowFunction(a *Assertion, name string, fn func(), expectExit bool) {
	isTerminated := internal.CheckTermination(fn)
	if expectExit {
		a.TrueNow(isTerminated, "%s() execution stopped, want do not stop", name)
	} else {
		a.NotTrueNow(isTerminated, "%s() execution do not stopped, want stop", name)
	}
}
