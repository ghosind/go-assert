package assert

import (
	"testing"
)

func TestPanicAndNotPanic(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	testPanicAndNotPanic(a, mockA, func() {
		// no panic
	}, false)
	testPanicAndNotPanic(a, mockA, func() {
		panic("some panic")
	}, true)
}

func testPanicAndNotPanic(a, mockA *Assertion, fn func(), isPanic bool) {
	// Panic
	testAssertionFunction(a, "Panic", func() error {
		return Panic(mockA.T, fn)
	}, isPanic)
	testAssertionFunction(a, "Assertion.Panic", func() error {
		return mockA.Panic(fn)
	}, isPanic)

	// NotPanic
	testAssertionFunction(a, "NotPanic", func() error {
		return NotPanic(mockA.T, fn)
	}, !isPanic)
	testAssertionFunction(a, "Assertion.NotPanic", func() error {
		return mockA.NotPanic(fn)
	}, !isPanic)

	// PanicNow
	testAssertionNowFunction(a, "PanicNow", func() {
		PanicNow(mockA.T, fn)
	}, !isPanic)
	testAssertionNowFunction(a, "Assertion.PanicNow", func() {
		mockA.PanicNow(fn)
	}, !isPanic)

	// NotPanicNow
	testAssertionNowFunction(a, "NotPanicNow", func() {
		NotPanicNow(mockA.T, fn)
	}, isPanic)
	testAssertionNowFunction(a, "Assertion.NotPanicNow", func() {
		mockA.NotPanicNow(fn)
	}, isPanic)
}

func TestIsPanic(t *testing.T) {
	Nil(t, isPanic(func() {
		// no panic
	}))
	NotNil(t, isPanic(func() {
		panic("unexpected panic")
	}))
}
