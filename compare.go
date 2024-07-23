package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// DeepEqual tests the deep equality between actual and expect parameters. It'll set the result to
// fail if they are not deeply equal, and it doesn't stop the execution.
//
//	a := assert.New(t)
//	a.DeepEqual(1, 1) // success
//	a.DeepEqual("ABC", "ABC") // success
//	a.DeepEqual(1, 0) // fail
//	a.DeepEqual(1, int64(1)) // fail
func (a *Assertion) DeepEqual(actual, expect any, message ...any) error {
	a.Helper()

	return tryDeepEqual(a.T, false, actual, expect, message...)
}

// DeepEqualNow tests the deep equality between actual and expect parameters, and it'll stop the
// execution if they are not deeply equal.
//
//	a := assert.New(t)
//	a.DeepEqualNow(1, 1) // success
//	a.DeepEqualNow("ABC", "ABC") // success
//	a.DeepEqualNow(1, int64(1)) // fail and terminate
//	// never run
func (a *Assertion) DeepEqualNow(actual, expect any, message ...any) error {
	a.Helper()

	return tryDeepEqual(a.T, true, actual, expect, message...)
}

// NotDeepEqual tests the deep inequality between actual and expected parameters. It'll set the
// result to fail if they are deeply equal, but it doesn't stop the execution.
//
//	a := assert.New(t)
//	a.NotDeepEqual(1, 0) // success
//	a.NotDeepEqual(1, int64(1)) // success
//	a.NotDeepEqual(1, 1) // fail
//	a.NotDeepEqual("ABC", "ABC") // fail
func (a *Assertion) NotDeepEqual(actual, expect any, message ...any) error {
	a.Helper()

	return tryNotDeepEqual(a.T, false, actual, expect, message...)
}

// NotDeepEqualNow tests the deep inequality between actual and expected parameters, and it'll stop
// the execution if they are deeply equal.
//
//	a := assert.New(t)
//	a.NotDeepEqual1, 0) // success
//	a.NotDeepEqual1, int64(1)) // success
//	a.NotDeepEqual"ABC", "ABC") // fail and terminate
//	// never run
func (a *Assertion) NotDeepEqualNow(actual, expect any, message ...any) error {
	a.Helper()

	return tryNotDeepEqual(a.T, true, actual, expect, message...)
}

// tryDeepEqual try to testing the deeply equality between actual and expect values, and it'll
// fail if the values are not deeply equal.
func tryDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return reflect.DeepEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageEqual, actual, expect),
		message...,
	)
}

// tryNotDeepEqual try to testing the deeply inequality between actual and expect values, and it'll
// fail if the values are deeply equal.
func tryNotDeepEqual(t *testing.T, failedNow bool, actual, expect any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return !reflect.DeepEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotEqual, actual, expect),
		message...,
	)
}

// Equal tests the equality between actual and expect parameters. It'll set the result to fail if
// they are not equal, and it doesn't stop the execution.
//
//	a := assert.New(t)
//	a.Equal(1, 1) // success
//	a.Equal("ABC", "ABC") // success
//	a.Equal(1, int64(1)) // success
//	a.Equal(1, uint64(1)) // fail
//	a.Equal(1, 0) // fail
func (a *Assertion) Equal(actual, expect any, message ...any) error {
	a.Helper()

	return tryEqual(a.T, false, actual, expect, message...)
}

// EqualNow tests the equality between actual and expect parameters, and it'll stop the execution
// if they are not equal.
//
//	a := assert.New(t)
//	a.EqualNow(1, 1) // success
//	a.EqualNow("ABC", "ABC") // success
//	a.EqualNow(1, int64(1)) // success
//	a.EqualNow(1, 0) // fail and terminate
//	never run
func (a *Assertion) EqualNow(actual, expect any, message ...any) error {
	a.Helper()

	return tryEqual(a.T, true, actual, expect, message...)
}

// NotEqual tests the inequality between actual and expected parameters. It'll set the result to
// fail if they are equal, but it doesn't stop the execution.
//
//	a := assert.New(t)
//	a.NotEqual(1, 0) // success
//	a.NotEqual("ABC", "CBA") // success
//	a.NotEqual(1, uint64(1)) // success
//	a.NotEqual(1, 1) // fail
//	a.NotEqual("ABC", "ABC") // fail
//	a.NotEqual(1, int64(1)) // fail
func (a *Assertion) NotEqual(actual, expect any, message ...any) error {
	a.Helper()

	return tryNotEqual(a.T, false, actual, expect, message...)
}

// NotEqualNow tests the inequality between actual and expected parameters, and it'll stop the
// execution if they are equal.
//
//	a := assert.New(t)
//	a.NotEqualNow(1, 0) // success
//	a.NotEqualNow("ABC", "CBA") // success
//	a.NotEqualNow(1, 1) // fail and terminate
//	// never run
func (a *Assertion) NotEqualNow(actual, expect any, message ...any) error {
	a.Helper()

	return tryNotEqual(a.T, true, actual, expect, message...)
}

// tryEqual try to testing the equality between actual and expect values, and it'll fail if the
// values are not equal.
func tryEqual(t *testing.T, failedNow bool, actual, expect any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return isEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageEqual, actual, expect),
		message...,
	)
}

// tryNotEqual try to testing the inequality between actual and expect values, and it'll fail if
// the values are equal.
func tryNotEqual(t *testing.T, failedNow bool, actual, expect any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isEqual(actual, expect) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotEqual, actual, expect),
		message...,
	)
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
//	a := assert.New(t)
//	var err error // nil
//	a.Nil(err) // success
//
//	err = errors.New("some error")
//	a.Nil(err) // fail
func (a *Assertion) Nil(val any, message ...any) error {
	a.Helper()

	return tryNil(a.T, false, val, message...)
}

// NilNow tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is not nil.
//
//	a := assert.New(t)
//	var err error // nil
//	a.NilNow(err) // success
//
//	err = errors.New("some error")
//	a.NilNow(err) // fail and terminate
//	// never run
func (a *Assertion) NilNow(val any, message ...any) error {
	a.Helper()

	return tryNil(a.T, true, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
//	a := assert.New(t)
//	var err error // nil
//	a.NotNil(err) // fail
//
//	err = errors.New("some error")
//	a.NotNil(err) // success
func (a *Assertion) NotNil(val any, message ...any) error {
	a.Helper()

	return tryNotNil(a.T, false, val, message...)
}

// NotNilNow tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is nil.
//
//	a := assert.New(t)
//	var err error = errors.New("some error")
//	a.NotNilNow(err) // success
//
//	err = nil
//	a.NotNilNow(err) // fail and terminate
//	// never run
func (a *Assertion) NotNilNow(val any, message ...any) error {
	a.Helper()

	return tryNotNil(a.T, true, val, message...)
}

// tryNil try to testing a value is nil or not, and it'll fail the value is nil.
func tryNil(t *testing.T, failedNow bool, val any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return isNil(val) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNil, val),
		message...,
	)
}

// tryNotNil try to testing a value is nil or not, and it'll fail the value is not nil.
func tryNotNil(t *testing.T, failedNow bool, val any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isNil(val) },
		failedNow,
		defaultErrMessageNotNil,
		message...,
	)
}

// True tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
//	a := assert.New(t)
//	a.True(1) // success
//	a.True("test") // success
//	a.True(0) // fail
//	a.True("") // fail
func (a *Assertion) True(val any, message ...any) error {
	a.Helper()

	return tryTrue(a.T, false, val, message...)
}

// TrueNow tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
// The function will stop the execution if the value is falsy.
//
//	a := assert.New(t)
//	a.TrueNow(1) // success
//	a.TrueNow("test") // success
//	a.TrueNow("") // fail and terminate
//	// never run
func (a *Assertion) TrueNow(val any, message ...any) error {
	a.Helper()

	return tryTrue(a.T, true, val, message...)
}

// NotTrue tests whether a value is truthy or not. It'll set the result to fail if the value is a
// truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
//	a := assert.New(t)
//	a.NotTrue(0) // success
//	a.NotTrue("") // success
//	a.NotTrue(1) // fail
//	a.NotTrue("test") // fail
func (a *Assertion) NotTrue(val any, message ...any) error {
	a.Helper()

	return tryNotTrue(a.T, false, val, message...)
}

// NotTrueNow tests whether a value is truthy or not. It'll set the result to fail if the value is
// a truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
// The function will stop the execution if the value is truthy.
//
//	a := assert.New(t)
//	a.NotTrueNow(0) // success
//	a.NotTrueNow("") // success
//	a.NotTrueNow("test") // fail and terminate
//	// never run
func (a *Assertion) NotTrueNow(val any, message ...any) error {
	a.Helper()

	return tryNotTrue(a.T, true, val, message...)
}

// tryTrue try to testing a value is truthy or falsy, and it'll fail the value is falsy.
func tryTrue(t *testing.T, failedNow bool, val any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool {
			return isTrue(val)
		},
		failedNow,
		defaultErrMessageTrue,
		message...,
	)
}

// tryNotTrue try to testing a value is truthy or falsy, and it'll fail the value is truthy.
func tryNotTrue(t *testing.T, failedNow bool, val any, message ...any) error {
	t.Helper()

	return test(
		t,
		func() bool {
			return !isTrue(val)
		},
		failedNow,
		defaultErrMessageNotTrue,
		message...,
	)
}

// isEqual checks the equality of the values.
func isEqual(x, y any) bool {
	if x == nil || y == nil {
		return x == y
	}

	var v1, v2 reflect.Value
	if xv, ok := x.(reflect.Value); ok {
		v1 = xv
	} else {
		v1 = reflect.ValueOf(x)
	}
	if yv, ok := y.(reflect.Value); ok {
		v2 = yv
	} else {
		v2 = reflect.ValueOf(y)
	}

	if ok := isSameType(v1.Type(), v2.Type()); !ok {
		return false
	}

	switch v1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Slice:
		return isSliceEqual(v1, v2)
	default:
		return v1.Interface() == v2.Interface()
	}
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
