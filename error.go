package assert

import "fmt"

// AssertionError indicates the failure of an assertion.
type AssertionError struct {
	message  *string
	actual   any
	expect   any
	operator string
}

func newAssertionError(operator string, actual, expect any, message ...string) AssertionError {
	err := AssertionError{
		actual:   actual,
		expect:   expect,
		operator: operator,
	}

	if len(message) > 0 {
		err.message = &message[0]
	}

	return err
}

func (err AssertionError) Error() string {
	if err.message != nil {
		return *err.message
	}

	return fmt.Sprintf("%v %s %v", err.actual, err.operator, err.expect)
}
