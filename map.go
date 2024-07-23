package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// MapHasKey tests whether the map contains the specified key or not, it will fail if the map does
// not contain the key, or the type of the key cannot assign to the type of the key of the map.
//
//	a := assert.New(t)
//	a.MapHasKey(map[string]int{"a":1}, "a") // success
//	a.MapHasKey(map[string]int{"a":1}, "b") // fail
//	a.MapHasKey(map[string]int{"a":1}, 1) // fail
func (a *Assertion) MapHasKey(m, key any, message ...any) error {
	a.Helper()

	return tryMapHasKey(a.T, false, m, key, message...)
}

// MapHasKeyNow tests whether the map contains the specified key or not, and it will terminate the
// execution if the test fails. It will fail if the map does not contain the key, or the type of
// the key cannot assign to the type of the key of the map.
//
//	a := assert.New(t)
//	a.MapHasKeyNow(map[string]int{"a":1}, "a") // success
//	a.MapHasKeyNow(map[string]int{"a":1}, "b") // fail and terminate
//	// never run
func (a *Assertion) MapHasKeyNow(m, key any, message ...any) error {
	a.Helper()

	return tryMapHasKey(a.T, true, m, key, message...)
}

// NotMapHasKey tests whether the map contains the specified key or not, it will fail if the map
// contain the key. It will also set the test result to success if the type of the key cannot
// assign to the type of the key of the map.
//
//	a := assert.New(t)
//	a.NotMapHasKey(map[string]int{"a":1}, "b") // success
//	a.NotMapHasKey(map[string]int{"a":1}, 1) // success
//	a.NotMapHasKey(map[string]int{"a":1}, "a") // fail
func (a *Assertion) NotMapHasKey(m, key any, message ...any) error {
	a.Helper()

	return tryNotMapHasKey(a.T, false, m, key, message...)
}

// NotMapHasKeyNow tests whether the map contains the specified key or not, it will fail if the map
// contain the key, and it will terminate the execution if the test fails. It will also set the
// test result to success if the type of the key cannot assign to the type of the key of the map.
//
//	a := assert.New(t)
//	a.NotMapHasKeyNow(map[string]int{"a":1}, "b") // success
//	a.NotMapHasKeyNow(map[string]int{"a":1}, 1) // success
//	a.NotMapHasKeyNow(map[string]int{"a":1}, "a") // fail and terminate
//	// never run
func (a *Assertion) NotMapHasKeyNow(m, key any, message ...any) error {
	a.Helper()

	return tryNotMapHasKey(a.T, true, m, key, message...)
}

// tryMapHasKey tries to test whether the map contains the specified key or not, and it'll fail if
// the map does not contains the specified key.
func tryMapHasKey(
	t *testing.T,
	failedNow bool,
	m, key any,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return isMapHasKey(m, key) },
		failedNow,
		fmt.Sprintf(defaultErrMessageMapHasKey, key),
		message...,
	)
}

// tryNotMapHasKey tries to test whether the map contains the specified key or not, and it'll fail
// if the map contains the specified key.
func tryNotMapHasKey(
	t *testing.T,
	failedNow bool,
	m, key any,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isMapHasKey(m, key) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotMapHasKey, key),
		message...,
	)
}

// MapHasValue tests whether the map contains the specified value or not, it will fail if the map
// does not contain the value, or the type of the value cannot assign to the type of the values of
// the map.
//
//	a := assert.New(t)
//	a.MapHasValue(map[string]int{"a":1}, 1) // success
//	a.MapHasValue(map[string]int{"a":1}, 2) // fail
//	a.MapHasValue(map[string]int{"a":1}, "a") // fail
func (a *Assertion) MapHasValue(m, value any, message ...any) error {
	a.Helper()

	return tryMapHasValue(a.T, false, m, value, message...)
}

// MapHasValueNow tests whether the map contains the specified value or not, and it will terminate
// the execution if the test fails. It will fail if the map does not contain the value, or the type
// of the value cannot assign to the type of the value of the map.
//
//	a := assert.New(t)
//	a.MapHasValueNow(map[string]int{"a":1}, 1) // success
//	a.MapHasValueNow(map[string]int{"a":1}, 2) // fail and terminate
//	// never run
func (a *Assertion) MapHasValueNow(m, value any, message ...any) error {
	a.Helper()

	return tryMapHasValue(a.T, true, m, value, message...)
}

// NotMapHasValue tests whether the map contains the specified value or not, it will fail if the
// map contain the value. It will also set the test result to success if the type of the value
// cannot assign to the type of the value of the map.
//
//	a := assert.New(t)
//	a.NotMapHasValue(map[string]int{"a":1}, 2) // success
//	a.NotMapHasValue(map[string]int{"a":1}, "a") // success
//	a.NotMapHasValue(map[string]int{"a":1}, 1) // fail
func (a *Assertion) NotMapHasValue(m, value any, message ...any) error {
	a.Helper()

	return tryNotMapHasValue(a.T, false, m, value, message...)
}

// NotMapHasValueNow tests whether the map contains the specified value or not, it will fail if the
// map contain the value, and it will terminate the execution if the test fails. It will also set
// the test result to success if the type of the value cannot assign to the type of the value of
// the map.
//
//	a := assert.New(t)
//	a.NotMapHasValueNow(map[string]int{"a":1}, 2) // success
//	a.NotMapHasValueNow(map[string]int{"a":1}, "a") // success
//	a.NotMapHasValueNow(map[string]int{"a":1}, 1) // fail and terminate
//	// never run
func (a *Assertion) NotMapHasValueNow(m, value any, message ...any) error {
	a.Helper()

	return tryNotMapHasValue(a.T, true, m, value, message...)
}

// tryMapHasValue tries to test whether the map contains the specified value or not, and it'll fail
// if the map does not contains the specified value.
func tryMapHasValue(
	t *testing.T,
	failedNow bool,
	m, value any,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return isMapHasValue(m, value) },
		failedNow,
		fmt.Sprintf(defaultErrMessageMapHasValue, value),
		message...,
	)
}

// tryNotMapHasValue tries to test whether the map contains the specified value or not, and it'll
// fail if the map contains the specified value.
func tryNotMapHasValue(
	t *testing.T,
	failedNow bool,
	m, value any,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isMapHasValue(m, value) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotMapHasValue, value),
		message...,
	)
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
