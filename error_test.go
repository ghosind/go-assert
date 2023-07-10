package assert

import "testing"

func TestAssertionError(t *testing.T) {
	err := newAssertionError("default message")
	DeepEqual(t, err.Error(), "assert error: default message")

	err = newAssertionError("default message", "custom message")
	DeepEqual(t, err.Error(), "custom message")
}
