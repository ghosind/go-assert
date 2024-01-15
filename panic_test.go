package assert

import (
	"errors"
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
	a.T.Helper()

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

func TestPanicOf(t *testing.T) {
	a := New(t)
	mockA := New(new(testing.T))

	expectedErr := errors.New("expected error")

	testPanicOf(a, mockA, func() {}, expectedErr, false)
	testPanicOf(a, mockA, func() {
		panic(expectedErr)
	}, expectedErr, true)
	testPanicOf(a, mockA, func() {
		panic("not expected error")
	}, expectedErr, false)
	testPanicOf(a, mockA, func() {
		panic("expected error")
	}, "expected error", true)
}

func testPanicOf(a, mockA *Assertion, fn func(), expectErr any, isExpectedPanic bool) {
	a.Helper()

	testAssertionFunction(a, "PanicOf", func() error {
		return PanicOf(mockA.T, fn, expectErr)
	}, isExpectedPanic)
	testAssertionFunction(a, "Assertion.PanicOf", func() error {
		return mockA.PanicOf(fn, expectErr)
	}, isExpectedPanic)

	testAssertionFunction(a, "NotPanicOf", func() error {
		return NotPanicOf(mockA.T, fn, expectErr)
	}, !isExpectedPanic)
	testAssertionFunction(a, "Assertion.NotPanicOf", func() error {
		return mockA.NotPanicOf(fn, expectErr)
	}, !isExpectedPanic)

	testAssertionNowFunction(a, "PanicOfNow", func() {
		PanicOfNow(mockA.T, fn, expectErr)
	}, !isExpectedPanic)
	testAssertionNowFunction(a, "Assertion.PanicOfNow", func() {
		mockA.PanicOfNow(fn, expectErr)
	}, !isExpectedPanic)

	testAssertionNowFunction(a, "NotPanicOfNow", func() {
		NotPanicOfNow(mockA.T, fn, expectErr)
	}, isExpectedPanic)
	testAssertionNowFunction(a, "Assertion.NotPanicOfNow", func() {
		mockA.NotPanicOfNow(fn, expectErr)
	}, isExpectedPanic)
}

func TestIsPanic(t *testing.T) {
	Nil(t, isPanic(func() {
		// no panic
	}))
	NotNil(t, isPanic(func() {
		panic("unexpected panic")
	}))
}
