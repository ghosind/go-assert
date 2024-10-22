package assert

import (
	"errors"
	"fmt"
)

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
	defaultErrMessagePanicOf            string = "expect panic by %v, got %v"
	defaultErrMessageNotPanicOf         string = "got unexpected panic error: %v"
	defaultErrMessageTrue               string = "the expression evaluated to a falsy value"
	defaultErrMessageNotTrue            string = "the expression evaluated to a truthy value"
	defaultErrMessageMapHasKey          string = "expect map has key %v"
	defaultErrMessageNotMapHasKey       string = "expect map has no key %v"
	defaultErrMessageMapHasValue        string = "expect map has value %v"
	defaultErrMessageNotMapHasValue     string = "expect map has no value %v"
	defaultErrMessageGt                 string = "%v must greater than %v"
	defaultErrMessageGte                string = "%v must greater than or equal to %v"
	defaultErrMessageLt                 string = "%v must less then %v"
	defaultErrMessageLte                string = "%v must less then or equal to %v"
	defaultErrMessageIsError            string = "expect err matches %v, got %v"
	defaultErrMessageNotIsError         string = "expect err does not matches %v"
)

var (
	// ErrNotArray indicates that the value must be a slice or an array.
	ErrNotArray error = errors.New("the value must be a slice or an array")
	// ErrNotFloat indicates that the value must be a floating number.
	ErrNotFloat error = errors.New("the value must be a floating number")
	// ErrNotMap indicates that the value must be a map.
	ErrNotMap error = errors.New("the value must be a map")
	// ErrNotOrderable indicates that the value must be orderable.
	ErrNotOrderable error = errors.New("the value must be orderable")
	// ErrNotSameType indicates that both values must be the same type.
	ErrNotSameType error = errors.New("the values must be the same type")
	// ErrRequireT indicates that the instance of testing.T is a required parameter.
	ErrRequireT error = errors.New("testing.T is required")
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
