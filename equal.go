package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// DeepEqual tests deeply equality between actual and expect parameters.
func DeepEqual(t *testing.T, actual, expect any, message ...string) error {
	if reflect.DeepEqual(actual, expect) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("%v == %v", actual, expect), message...)

	t.Error(err)

	return err
}

// NotDeepEqual tests deeply inequality between actual and expected parameters.
func NotDeepEqual(t *testing.T, actual, expect any, message ...string) error {
	if !reflect.DeepEqual(actual, expect) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("%v != %v", actual, expect), message...)

	t.Error(err)

	return err
}

// Nil tests a value is nil or not, and it'll failed when the value is not nil.
func Nil(t *testing.T, val any, message ...string) error {
	if isNil(val) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("expect nil, got %v", val), message...)

	t.Error(err)

	return err
}

// NotNil tests a value is nil or not, and it'll failed when the value is nil.
func NotNil(t *testing.T, val any, message ...string) error {
	if !isNil(val) {
		return nil
	}

	err := newAssertionError("expect not nil, got nil", message...)

	t.Error(err)

	return err
}
