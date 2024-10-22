package assert

import (
	"testing"

	"github.com/ghosind/go-assert/internal"
)

func TestFailedHandler(t *testing.T) {
	mockT := new(testing.T)
	assert := New(t)

	failed(mockT, nil, false)
	assert.NotTrue(mockT.Failed(), false)

	failed(mockT, newAssertionError("Test error"), false)
	assert.True(mockT.Failed())

	isTerminated := internal.CheckTermination(func() {
		failed(mockT, newAssertionError("Test error"), true)
	})
	assert.True(isTerminated)
}

func TestToFloat(t *testing.T) {
	a := New(t)

	a.NotPanicNow(func() {
		a.Equal(toFloat(1), 1.0)
	})
	a.NotPanicNow(func() {
		a.Equal(toFloat(1.0), 1.0)
	})
	a.PanicNow(func() {
		toFloat("1.0")
	})
}
