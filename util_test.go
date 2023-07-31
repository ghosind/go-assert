package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestFailedHandler(t *testing.T) {
	mockT := new(testing.T)
	assert := New(t)

	failed(mockT, nil, false)
	assert.DeepEqual(mockT.Failed(), false)

	failed(mockT, newAssertionError("Test error"), false)
	assert.DeepEqual(mockT.Failed(), true)

	isTerminated := internal.CheckTermination(func() {
		failed(mockT, newAssertionError("Test error"), true)
	})
	assert.DeepEqual(isTerminated, true)
}

func TestIsNil(t *testing.T) {
	assert := New(t)

	assert.DeepEqual(isNil(1), false)  // int
	assert.DeepEqual(isNil(""), false) // string
	assert.DeepEqual(isNil(nil), true)
	var testAssert *Assertion
	assert.DeepEqual(isNil(testAssert), true)
	assert.DeepEqual(isNil(assert), false)
}

func TestIsPanic(t *testing.T) {
	Nil(t, isPanic(func() {
		// no panic
	}))
	NotNil(t, isPanic(func() {
		panic("unexpected panic")
	}))
}
