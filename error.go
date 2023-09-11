package assert

import "fmt"

const (
	defaultErrMessageEqual              string = "%v == %v"
	defaultErrMessageNotEqual           string = "%v != %v"
	defaultErrMessageContainsElement    string = "expect contains %v"
	defaultErrMessageNotContainsElement string = "expect did not contains %v"
	defaultErrMessageContainsString     string = "expect contains \"%s\""
	defaultErrMessageNotContainsString  string = "expect did not contain \"%s\""
	defaultErrMessageHasPrefixString    string = "expect has prefix \"%s\""
	defaultErrMessageNotHasPrefixString string = "expect has no prefix \"%s\""
	defaultErrMessageHasSuffixString    string = "expect has suffix \"%s\""
	defaultErrMessageNotHasSuffixString string = "expect has no suffix \"%s\""
	defaultErrMessageMatch              string = "the input did not match the regular expression"
	defaultErrMessageNotMatch           string = "the input match the regular expression"
	defaultErrMessageNil                string = "expect nil, got %v"
	defaultErrMessageNotNil             string = "expect not nil, got nil"
	defaultErrMessagePanic              string = "missing expected panic"
	defaultErrMessageNotPanic           string = "got unwanted error: %v"
	defaultErrMessageTrue               string = "the expression evaluated to a falsy value"
	defaultErrMessageNotTrue            string = "the expression evaluated to a truthy value"
)

// AssertionError indicates the failure of an assertion.
type AssertionError struct {
	message string
}

// newAssertionError creates a new error with custom message or default message.
func newAssertionError(defaultMsg string, message ...any) AssertionError {
	err := AssertionError{}

	if len(message) > 0 {
		if format, ok := message[0].(string); ok {
			err.message = fmt.Sprintf(format, message[1:]...)
		}
	}

	if err.message == "" {
		err.message = "assert error: " + defaultMsg
	}

	return err
}

// Error returns the message of the error.
func (err AssertionError) Error() string {
	return err.message
}
