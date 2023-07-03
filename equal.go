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
