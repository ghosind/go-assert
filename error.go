package assert

// AssertionError indicates the failure of an assertion.
type AssertionError struct {
	message string
}

// newAssertionError creates a new error with custom message or default message.
func newAssertionError(defaultMsg string, message ...string) AssertionError {
	err := AssertionError{}

	if len(message) > 0 {
		err.message = message[0]
	} else {
		err.message = "assert error: " + defaultMsg
	}

	return err
}

// Error returns the message of the error.
func (err AssertionError) Error() string {
	return err.message
}
