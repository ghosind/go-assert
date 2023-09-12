package assert

import "testing"

func TestAssertionError(t *testing.T) {
	err := newAssertionError("default message")
	Equal(t, err.Error(), "assert error: default message")

	err = newAssertionError("default message", "custom message")
	Equal(t, err.Error(), "custom message")

	err = newAssertionError("default message", "custom message with argument: %v", 1)
	Equal(t, err.Error(), "custom message with argument: 1")

	err = newAssertionError("default message", 1)
	Equal(t, err.Error(), "assert error: default message")
}
