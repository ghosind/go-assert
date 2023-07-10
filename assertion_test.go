package assert

import (
	"os"
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

func TestFailed(t *testing.T) {
	testFailed(t, func(a *Assertion) {
		// do nothing
	}, false)
	testFailed(t, func(a *Assertion) {
		a.Log("test")
	}, false)
	testFailed(t, func(a *Assertion) {
		a.Logf("test")
	}, false)
	testFailed(t, func(a *Assertion) {
		a.Error("test")
	}, true)
	testFailed(t, func(a *Assertion) {
		a.Errorf("test")
	}, true)
	testFailed(t, func(a *Assertion) {
		a.Fail()
	}, true)
}

func testFailed(t *testing.T, f func(a *Assertion), failed bool) {
	assert := New(new(testing.T))

	f(assert)
	DeepEqual(t, assert.Failed(), failed)
}

func TestRun(t *testing.T) {
	assert := New(t)

	assert.Setenv("env", "assert test")
	assert.Run("test Run", func(t *testing.T) {
		DeepEqual(t, t.Name(), "TestRun/test_Run")
		DeepEqual(t, os.Getenv("env"), "assert test")
	})
}
