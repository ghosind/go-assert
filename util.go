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

// isEqual checks the equality of the values.
func isEqual(x, y any) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if !isSameType(v1.Type(), v2.Type()) {
		return false
	}

	switch v1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.Slice:
		return isSliceEqual(v1, v2)
	default:
		return x == y
	}
}

// isSameType indicates the equality of two types, and it will ignore the bit size of the same
// type. For example, `int32` and `int64` will be the same type.
func isSameType(t1, t2 reflect.Type) bool {
	kind := t2.Kind()

	switch t1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return kind >= reflect.Int && kind <= reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return kind >= reflect.Uint && kind <= reflect.Uint64
	case reflect.Float32, reflect.Float64:
		return kind == reflect.Float32 || kind == reflect.Float64
	case reflect.Complex64, reflect.Complex128:
		return kind == reflect.Complex64 || kind == reflect.Complex128
	default:
		return t1 == t2
	}
}

// isSliceEqual checks the equality of each elements in the slices.
func isSliceEqual(v1, v2 reflect.Value) bool {
	if v1.Len() != v2.Len() {
		return false
	}

	for i := 0; i < v1.Len(); i++ {
		if v1.Index(i).Interface() != v2.Index(i).Interface() {
			return false
		}
	}

	return true
}

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
