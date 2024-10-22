package assert

import (
	"reflect"
	"testing"
)

var (
	floatType = reflect.TypeOf(float64(0))
)

// test tries to run the test function, and creates an assertion error if the result is fail.
func test(
	t *testing.T,
	fn func() bool,
	failedNow bool,
	defaultMessage string,
	message ...any,
) error {
	t.Helper()

	if fn() {
		return nil
	}

	err := newAssertionError(defaultMessage, message...)

	failed(t, err, failedNow)

	return err
}

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

// isSameType indicates the equality of two types, and it will ignore the bit size of the same
// type. For example, `int32` and `int64` will be the same type.
//
// It will returns a bool value to tell the callee function that one of the values is an integer,
// but another one is unsigned. For this case, it needs to check the value to compare them.
func isSameType(t1, t2 reflect.Type) bool {
	kind := t2.Kind()

	switch t1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return kind >= reflect.Int && kind <= reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return kind >= reflect.Uint && kind <= reflect.Uintptr
	case reflect.Float32, reflect.Float64:
		return kind == reflect.Float32 || kind == reflect.Float64
	case reflect.Complex64, reflect.Complex128:
		return kind == reflect.Complex64 || kind == reflect.Complex128
	default:
		return t1 == t2
	}
}

// toFloat converts the value to a float64, and it'll panic if the value can't be converted.
func toFloat(v any) float64 {
	vv := reflect.ValueOf(v)
	if vv.CanFloat() {
		return vv.Float()
	}

	if vv.CanConvert(floatType) {
		return vv.Convert(floatType).Float()
	}

	panic(ErrNotFloat)
}
