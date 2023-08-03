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
