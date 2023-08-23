package assert

import "testing"

func TestAssertionError(t *testing.T) {
	err := newAssertionError("default message")
	DeepEqual(t, err.Error(), "assert error: default message")

	err = newAssertionError("default message", "custom message")
	DeepEqual(t, err.Error(), "custom message")

	err = newAssertionError("default message", "custom message with argument: %v", 1)
	DeepEqual(t, err.Error(), "custom message with argument: 1")

	err = newAssertionError("default message", 1)
	DeepEqual(t, err.Error(), "assert error: default message")
}
