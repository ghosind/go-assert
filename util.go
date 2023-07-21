package assert

import (
	"reflect"
	"testing"
)

// failed handles the assertion error with the specific testing.T or the assertion's t. It will set
// marks the function has failed if the err is not nil. It'll also stops the execution if failedNow
// set to true.
func failed(t *testing.T, err error, failedNow bool) {
	if err == nil {
		return
	}

	if failedNow {
		t.Fatal(err)
	} else {
		t.Error(err)
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
