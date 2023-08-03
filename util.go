package assert

import (
	"reflect"
	"testing"
)

// failed handles the assertion error with the specific testing.T or the assertion's t. It will set
// marks the function has failed if the err is not nil. It'll also stops the execution if failedNow
// set to true.
func failed(t *testing.T, err error, failedNow bool) {
	t.Helper()

	if err == nil {
		return
	}

	t.Error(err)

	if failedNow {
		t.FailNow()
	}
}

// ################################
// ## Assertion Helper Functions ##
// ################################

// isNil checks whether a value is nil or not. It'll always return false if the value is not a
// channel, a function, a map, a point, an unsafe point, an interface, or a slice.
func isNil(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer,
		reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// isPanic executes the function, and tries to catching and returns the return value from
// recover().
func isPanic(fn func()) (err any) {
	defer func() {
		if e := recover(); e != nil {
			err = e
		}
	}()

	fn()

	return
}

// isTrue checks whether a value is truthy or not. It'll return true if the value is not the zero
// value for its type. For a slice, a truthy value should not be the zero value and the length must
// be greater than 0. For nil, it'll always return false.
func isTrue(v any) bool {
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Invalid:
		return false // always false
	case reflect.Slice:
		return v != nil && rv.Len() > 0
	default:
		return !rv.IsZero()
	}
}
