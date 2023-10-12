package assert

import (
	"math"
	"reflect"
	"testing"
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

// isComparable gets the type of the value, and checks whether the type is comparable or not.
func isComparable(v any) bool {
	switch v.(type) {
	case
		int, int8, int16, int32, int64, // Signed integer
		uint, uint8, uint16, uint32, uint64, uintptr, // Unsigned integer
		float32, float64, // Floating-point number
		string: // string
		return true
	default:
		return false
	}
}

// isContainsElement checks whether the array or slice contains the specific element or not. It'll
// panic if the source is not an array or a slice, and it'll also panic if the element's type is
// not the same as the source's element.
func isContainsElement(source, elem any) bool {
	st := reflect.ValueOf(source)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	if st.Kind() != reflect.Array && st.Kind() != reflect.Slice {
		panic("require array or slice")
	}
	if ok, isMixed := isSameType(st.Type().Elem(), reflect.TypeOf(elem)); !ok && !isMixed {
		panic("require same type")
	}

	if st.Len() == 0 {
		return false
	}

	ev := reflect.ValueOf(elem)

	for i := 0; i < st.Len(); i++ {
		ok := isEqual(st.Index(i), ev)
		if ok {
			return true
		}
	}
	return false
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

	if isSame, isMixSign := isSameType(v1.Type(), v2.Type()); !isSame {
		if isMixSign {
			return isEqualForMixSignInt(v1, v2)
		}
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

// isEqualForMixSignInt checks the equality of two integers one of an integer is signed, but
// another one is unsigned.
func isEqualForMixSignInt(v1, v2 reflect.Value) bool {
	intVal := v1
	uintVal := v2
	if v1.Kind() >= reflect.Uint && v1.Kind() <= reflect.Uintptr {
		intVal = v2
		uintVal = v1
	}

	if intVal.Int() < 0 {
		return false
	} else if uintVal.Uint() > uint64(math.MaxInt64) {
		return false
	}

	return intVal.Int() == int64(uintVal.Uint())
}

// isMapHasValue checks whether the map contains the specified key or not.
func isMapHasKey(m, k any) bool {
	if m == nil || reflect.TypeOf(m).Kind() != reflect.Map {
		return false
	}

	mv := reflect.ValueOf(m)
	if mv.Len() == 0 {
		return false
	}

	if !reflect.TypeOf(k).AssignableTo(mv.Type().Key()) {
		return false
	}

	return mv.MapIndex(reflect.ValueOf(k)).Kind() != reflect.Invalid
}

// isMapHasValue checks whether the map contains the specified value or not.
func isMapHasValue(m, v any) bool {
	if m == nil || reflect.TypeOf(m).Kind() != reflect.Map {
		return false
	}

	mv := reflect.ValueOf(m)
	if mv.Len() == 0 {
		return false
	}

	if !reflect.TypeOf(v).AssignableTo(mv.Type().Elem()) {
		return false
	}

	vv := reflect.ValueOf(v)
	iter := mv.MapRange()

	for iter.Next() {
		mvv := iter.Value()
		if isEqual(mvv, vv.Convert(mvv.Type())) {
			return true
		}
	}

	return false
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

// isSameType indicates the equality of two types, and it will ignore the bit size of the same
// type. For example, `int32` and `int64` will be the same type.
//
// It will returns a bool value to tell the callee function that one of the values is an integer,
// but another one is unsigned. For this case, it needs to check the value to compare them.
func isSameType(t1, t2 reflect.Type) (isSame bool, isMixSign bool) {
	kind := t2.Kind()

	switch t1.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return kind >= reflect.Int && kind <= reflect.Int64,
			kind >= reflect.Uint && kind <= reflect.Uintptr
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return kind >= reflect.Uint && kind <= reflect.Uintptr,
			kind >= reflect.Int && kind <= reflect.Int64
	case reflect.Float32, reflect.Float64:
		return kind == reflect.Float32 || kind == reflect.Float64, false
	case reflect.Complex64, reflect.Complex128:
		return kind == reflect.Complex64 || kind == reflect.Complex128, false
	default:
		return t1 == t2, false
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
