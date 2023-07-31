package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestPanicAndNotPanic(t *testing.T) {
	mockT := new(testing.T)
	assertion := New(mockT)

	testPanicAndNotPanic(t, assertion, func() {
		// no panic
	}, false)
	testPanicAndNotPanic(t, assertion, func() {
		panic("some panic")
	}, true)
}

func testPanicAndNotPanic(t *testing.T, assertion *Assertion, fn func(), isPanic bool) {
	testPanic(t, assertion, fn, isPanic)

	testNotPanic(t, assertion, fn, isPanic)

	testPanicNow(t, assertion, fn, isPanic)

	testNotPanicNow(t, assertion, fn, isPanic)
}

func testPanic(t *testing.T, assertion *Assertion, fn func(), isPanic bool) {
	err := assertion.Panic(fn)
	if isPanic && err != nil {
		t.Errorf("Panic() = %v, want = nil", err)
	} else if !isPanic && err == nil {
		t.Error("Panic() = nil, want error")
	}

	err = Panic(assertion.T, fn)
	if isPanic && err != nil {
		t.Errorf("Panic() = %v, want = nil", err)
	} else if !isPanic && err == nil {
		t.Error("Panic() = nil, want error")
	}
}

func testNotPanic(t *testing.T, assertion *Assertion, fn func(), isPanic bool) {
	err := assertion.NotPanic(fn)
	if !isPanic && err != nil {
		t.Errorf("NotPanic() = %v, want = nil", err)
	} else if isPanic && err == nil {
		t.Error("NotPanic() = nil, want error")
	}

	err = NotPanic(assertion.T, fn)
	if !isPanic && err != nil {
		t.Errorf("NotPanic() = %v, want = nil", err)
	} else if isPanic && err == nil {
		t.Error("NotPanic() = nil, want error")
	}
}

func testPanicNow(t *testing.T, assertion *Assertion, fn func(), isPanic bool) {
	isTerminated := internal.CheckTermination(func() {
		assertion.PanicNow(fn)
	})
	if isPanic && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isPanic && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		PanicNow(assertion.T, fn)
	})
	if isPanic && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if !isPanic && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}

func testNotPanicNow(t *testing.T, assertion *Assertion, fn func(), isPanic bool) {
	isTerminated := internal.CheckTermination(func() {
		assertion.NotPanicNow(fn)
	})
	if !isPanic && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isPanic && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}

	isTerminated = internal.CheckTermination(func() {
		NotPanicNow(assertion.T, fn)
	})
	if !isPanic && isTerminated {
		t.Error("execution stopped, want do not stop")
	} else if isPanic && !isTerminated {
		t.Error("execution do not stopped, want stop")
	}
}
