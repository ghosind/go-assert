package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// Equal tests equality between actual and expect parameters.
func Equal(t *testing.T, actual, expect any, message ...string) error {
	if reflect.DeepEqual(actual, expect) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("%v == %v", actual, expect), message...)

	t.Error(err)

	return err
}

// NotEqual tests inequality between actual and expected parameters.
func NotEqual(t *testing.T, actual, expect any, message ...string) error {
	if !reflect.DeepEqual(actual, expect) {
		return nil
	}

	err := newAssertionError(fmt.Sprintf("%v != %v", actual, expect), message...)

	t.Error(err)

	return err
}
