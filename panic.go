package assert

import (
	"fmt"
	"testing"
)

// Panic expects the function fn to panic.
func Panic(t *testing.T, fn func(), message ...string) (err error) {
	defer func() {
		if e := recover(); e == nil {
			err = newAssertionError("mssing expected panic", message...)
		}
	}()

	fn()

	// err = newAssertionError("mssing expected panic", message...)

	return
}

// NotPanic asserts that the function fn does not panic.
func NotPanic(t *testing.T, fn func(), message ...string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = newAssertionError(fmt.Sprintf("got unwanted error: %v", e), message...)
		}
	}()

	fn()

	return
}
