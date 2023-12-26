package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// NotContainsElement tests whether the array or slice contains the specified element or not, and
// it set the result to fail if the array or slice does not contain the specified element. It'll
// panic if the `source` is not an array or a slice.
//
//	assertion.ContainsElement([]int{1, 2, 3}, 1) // success
//	assertion.ContainsElement([]int{1, 2, 3}, 3) // success
//	assertion.ContainsElement([]int{1, 2, 3}, 4) // fail
func (a *Assertion) ContainsElement(source, expect any, message ...any) error {
	a.Helper()

	return tryContainsElement(a.T, false, source, expect, message...)
}

// ContainsElementNow tests whether the array or slice contains the specified element or not, and
// it will terminate the execution if the array or slice does not contain the specified element.
// It'll panic if the `source` is not an array or a slice.
//
//	assertion.ContainsElementNow([]int{1, 2, 3}, 1) // success
//	assertion.ContainsElementNow([]int{1, 2, 3}, 3) // success
//	assertion.ContainsElementNow([]int{1, 2, 3}, 4) // fail and stop the execution
//	// never runs
func (a *Assertion) ContainsElementNow(source, expect any, message ...any) error {
	a.Helper()

	return tryContainsElement(a.T, true, source, expect, message...)
}

// NotContainsElement tests whether the array or slice contains the specified element or not, and
// it set the result to fail if the array or slice contains the specified element. It'll panic if
// the `source` is not an array or a slice.
//
//	assertion.NotContainsElement([]int{1, 2, 3}, 4) // success
//	assertion.NotContainsElement([]int{1, 2, 3}, 0) // success
//	assertion.NotContainsElement([]int{1, 2, 3}, 1) // fail
func (a *Assertion) NotContainsElement(source, expect any, message ...any) error {
	a.Helper()

	return tryNotContainsElement(a.T, false, source, expect, message...)
}

// NotContainsElementNow tests whether the array or slice contains the specified element or not,
// and it will terminate the execution if the array or slice contains the specified element. It'll
// panic if the `source` is not an array or a slice.
//
//	assertion.NotContainsElementNow([]int{1, 2, 3}, 4) // success
//	assertion.NotContainsElementNow([]int{1, 2, 3}, 0) // success
//	assertion.NotContainsElementNow([]int{1, 2, 3}, 1) // fail and stop the execution
//	// never runs
func (a *Assertion) NotContainsElementNow(source, expect any, message ...any) error {
	a.Helper()

	return tryNotContainsElement(a.T, true, source, expect, message...)
}

// tryContainsElement tries to test whether the array or slice contains the specified element or
// not, and it'll fail if the array or slice does not contains the specified element.
func tryContainsElement(
	t *testing.T,
	failedNow bool,
	src, elem any,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return isContainsElement(src, elem) },
		failedNow,
		fmt.Sprintf(defaultErrMessageContainsElement, elem),
		message...,
	)
}

// tryNotContainsElement tries to test whether the array or slice contains the specified element
// or not, and it'll fail if the array of slice contains the specified element.
func tryNotContainsElement(
	t *testing.T,
	failedNow bool,
	src, elem any,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return !isContainsElement(src, elem) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotContainsElement, elem),
		message...,
	)
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
		panic(ErrNotArray)
	}
	if ok := isSameType(st.Type().Elem(), reflect.TypeOf(elem)); !ok {
		panic(ErrNotSameType)
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
