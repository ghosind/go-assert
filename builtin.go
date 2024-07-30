package assert

import (
	"fmt"
	"regexp"
	"testing"
)

// NotContainsElement tests whether the array or slice contains the specified element or not, and
// it set the result to fail if the array or slice does not contain the specified element. It'll
// panic if the `source` is not an array or a slice.
//
//	assert.ContainsElement(t, []int{1, 2, 3}, 1) // success
//	assert.ContainsElement(t, []int{1, 2, 3}, 3) // success
//	assert.ContainsElement(t, []int{1, 2, 3}, 4) // fail
func ContainsElement(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryContainsElement(t, false, source, expect, message...)
}

// ContainsElementNow tests whether the array or slice contains the specified element or not, and
// it will terminate the execution if the array or slice does not contain the specified element.
// It'll panic if the `source` is not an array or a slice.
//
//	assert.ContainsElementNow(t, []int{1, 2, 3}, 1) // success
//	assert.ContainsElementNow(t, []int{1, 2, 3}, 3) // success
//	assert.ContainsElementNow(t, []int{1, 2, 3}, 4) // fail and stop the execution
//	// never runs
func ContainsElementNow(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryContainsElement(t, true, source, expect, message...)
}

// NotContainsElement tests whether the array or slice contains the specified element or not, and
// it set the result to fail if the array or slice contains the specified element. It'll panic if
// the `source` is not an array or a slice.
//
//	assert.NotContainsElement(t, []int{1, 2, 3}, 4) // success
//	assert.NotContainsElement(t, []int{1, 2, 3}, 0) // success
//	assert.NotContainsElement(t, []int{1, 2, 3}, 1) // fail
func NotContainsElement(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryNotContainsElement(t, false, source, expect, message...)
}

// NotContainsElementNow tests whether the array or slice contains the specified element or not,
// and it will terminate the execution if the array or slice contains the specified element. It'll
// panic if the `source` is not an array or a slice.
//
//	assert.NotContainsElementNow(t, []int{1, 2, 3}, 4) // success
//	assert.NotContainsElementNow(t, []int{1, 2, 3}, 0) // success
//	assert.NotContainsElementNow(t, []int{1, 2, 3}, 1) // fail and stop the execution
//	// never runs
func NotContainsElementNow(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryNotContainsElement(t, true, source, expect, message...)
}

// ContainsString tests whether the string contains the substring or not, and it set the result to
// fail if the string does not contains the substring.
//
//	assert.ContainsString(t, "Hello world", "") // success
//	assert.ContainsString(t, "Hello world", "Hello") // success
//	assert.ContainsString(t, "Hello world", "world") // success
//	assert.ContainsString(t, "Hello world", "hello") // fail
func ContainsString(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryContainsString(t, false, str, substr, message...)
}

// ContainsStringNow tests whether the string contains the substring or not, and it will terminate the
// execution if the string does not contains the substring.
//
//	assert.ContainsStringNow(t, "Hello world", "") // success
//	assert.ContainsStringNow(t, "Hello world", "Hello") // success
//	assert.ContainsStringNow(t, "Hello world", "world") // success
//	assert.ContainsStringNow(t, "Hello world", "hello") // fail and stop the execution
//	// never runs
func ContainsStringNow(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryContainsString(t, true, str, substr, message...)
}

// Gt compares the values and sets the result to false if the first value is not greater than to
// the second value.
//
//	assert.Gt(t, 2, 1) // success
//	assert.Gt(t, 3.14, 1.68) // success
//	assert.Gt(t, "BCD", "ABC") // success
//	assert.Gt(t, 2, 2) // fail
//	assert.Gt(t, 1, 2) // fail
func Gt(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		false,
		compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGt, v1, v2),
		message...,
	)
}

// GtNow compares the values and sets the result to false if the first value is not greater than to
// the second value. It will panic if they do not match the expected result.
//
//	assert.GtNow(t, 2, 1) // success
//	assert.GtNow(t, 3.14, 1.68) // success
//	assert.GtNow(t, "BCD", "ABC") // success
//	assert.GtNow(t, 1, 2) // fail and terminate
//	// never runs
func GtNow(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		true,
		compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGt, v1, v2),
		message...,
	)
}

// Gte compares the values and sets the result to false if the first value is not greater than or
// equal to the second value.
//
//	assert.Gte(t, 2, 1) // success
//	assert.Gte(t, 3.14, 1.68) // success
//	assert.Gte(t, "BCD", "ABC") // success
//	assert.Gte(t, 2, 2) // success
//	assert.Gte(t, 1, 2) // fail
func Gte(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		false,
		compareTypeEqual|compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGte, v1, v2),
		message...,
	)
}

// GteNow compares the values and sets the result to false if the first value is not greater than
// or equal to the second value. It will panic if they do not match the expected result.
//
//	assert.GteNow(t, 2, 1) // success
//	assert.GteNow(t, 3.14, 1.68) // success
//	assert.GteNow(t, "BCD", "ABC") // success
//	assert.GteNow(t, 2, 2) // success
//	assert.GteNow(t, 1, 2) // fail and terminate
//	// never runs
func GteNow(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		true,
		compareTypeEqual|compareTypeGreater,
		v1, v2,
		fmt.Sprintf(defaultErrMessageGte, v1, v2),
		message...,
	)
}

// Lt compares the values and sets the result to false if the first value is not less than the
// second value.
//
//	assert.Lt(t, 1, 2) // success
//	assert.Lt(t, 1.68, 3.14) // success
//	assert.Lt(t, "ABC", "BCD") // success
//	assert.Lt(t, 2, 2) // fail
//	assert.Lt(t, 2, 1) // fail
func Lt(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		false,
		compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLt, v1, v2),
		message...,
	)
}

// LtNow compares the values and sets the result to false if the first value is not less than the
// second value. It will panic if they do not match the expected result.
//
//	assert.LtNow(t, 1, 2) // success
//	assert.LtNow(t, 1.68, 3.14) // success
//	assert.LtNow(t, "ABC", "BCD") // success
//	assert.LtNow(t, 2, 1) // fail and terminate
//	// never runs
func LtNow(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		true,
		compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLt, v1, v2),
		message...,
	)
}

// Lte compares the values and sets the result to false if the first value is not less than or
// equal to the second value.
//
//	assert.Lte(t, 1, 2) // success
//	assert.Lte(t, 1.68, 3.14) // success
//	assert.Lte(t, "ABC", "BCD") // success
//	assert.Lte(t, 2, 2) // success
//	assert.Lte(t, 2, 1) // fail
func Lte(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		false,
		compareTypeEqual|compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLte, v1, v2),
		message...,
	)
}

// LteNow compares the values and sets the result to false if the first value is not less than or
// equal to the second value. It will panic if they do not match the expected result.
//
//	assert.LteNow(t, 1, 2) // success
//	assert.LteNow(t, 1.68, 3.14) // success
//	assert.LteNow(t, "ABC", "BCD") // success
//	assert.LteNow(t, 2, 2) // success
//	assert.LteNow(t, 2, 1) // fail and terminate
//	// never runs
func LteNow(t *testing.T, v1, v2 any, message ...string) error {
	t.Helper()

	return tryCompareOrderableValues(
		t,
		true,
		compareTypeEqual|compareTypeLess,
		v1, v2,
		fmt.Sprintf(defaultErrMessageLte, v1, v2),
		message...,
	)
}

// NotContainsString tests whether the string contains the substring or not, and it set the result
// to fail if the string contains the substring.
//
//	assert.NotContainsString(t, "Hello world", "") // fail
//	assert.NotContainsString(t, "Hello world", "Hello") // fail
//	assert.NotContainsString(t, "Hello world", "world") // fail
//	assert.NotContainsString(t, "Hello world", "hello") // success
func NotContainsString(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryNotContainsString(t, false, str, substr, message...)
}

// NotContainsStringNow tests whether the string contains the substring or not, and it will terminate the
// execution if the string contains the substring.
//
//	assert.NotContainsStringNow(t, "Hello world", "hello") // success
//	assert.NotContainsStringNow(t, "Hello world", "Hello") // fail and stop the execution
//	// never runs
func NotContainsStringNow(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryNotContainsString(t, true, str, substr, message...)
}

// DeepEqual tests the deep equality between actual and expect parameters. It'll set the result to
// fail if they are not deeply equal, and it doesn't stop the execution.
//
//	assert.DeepEqual(t, 1, 1) // success
//	assert.DeepEqual(t, "ABC", "ABC") // success
//	assert.DeepEqual(t, 1, 0) // fail
//	assert.DeepEqual(t, 1, int64(1)) // fail
func DeepEqual(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryDeepEqual(t, false, actual, expect, message...)
}

// DeepEqualNow tests the deep equality between actual and expect parameters, and it'll stop the
// execution if they are not deeply equal.
//
//	assert.DeepEqualNow(t, 1, 1) // success
//	assert.DeepEqualNow(t, "ABC", "ABC") // success
//	assert.DeepEqualNow(t, 1, int64(1)) // fail and terminate
//	// never run
func DeepEqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryDeepEqual(t, true, actual, expect, message...)
}

// NotDeepEqual tests the deep inequality between actual and expected parameters. It'll set the
// result to fail if they are deeply equal, but it doesn't stop the execution.
//
//	assert.NotDeepEqual(t, 1, 0) // success
//	assert.NotDeepEqual(t, 1, int64(1)) // success
//	assert.NotDeepEqual(t, 1, 1) // fail
//	assert.NotDeepEqual(t, "ABC", "ABC") // fail
func NotDeepEqual(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotDeepEqual(t, false, actual, expect, message...)
}

// NotDeepEqualNow tests the deep inequality between actual and expected parameters, and it'll stop
// the execution if they are deeply equal.
//
//	assert.NotDeepEqual(t, 1, 0) // success
//	assert.NotDeepEqual(t, 1, int64(1)) // success
//	assert.NotDeepEqual(t, "ABC", "ABC") // fail and terminate
//	// never run
func NotDeepEqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotDeepEqual(t, true, actual, expect, message...)
}

// Equal tests the equality between actual and expect parameters. It'll set the result to fail if
// they are not equal, and it doesn't stop the execution.
//
//	assert.Equal(t, 1, 1) // success
//	assert.Equal(t, "ABC", "ABC") // success
//	assert.Equal(t, 1, int64(1)) // success
//	assert.Equal(t, 1, uint64(1)) // fail
//	assert.Equal(t, 1, 0) // fail
func Equal(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryEqual(t, false, actual, expect, message...)
}

// EqualNow tests the equality between actual and expect parameters, and it'll stop the execution
// if they are not equal.
//
//	assert.EqualNow(t, 1, 1) // success
//	assert.EqualNow(t, "ABC", "ABC") // success
//	assert.EqualNow(t, 1, int64(1)) // success
//	assert.EqualNow(t, 1, 0) // fail and terminate
//	never run
func EqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryEqual(t, true, actual, expect, message...)
}

// NotEqual tests the inequality between actual and expected parameters. It'll set the result to
// fail if they are equal, but it doesn't stop the execution.
//
//	assert.NotEqual(t, 1, 0) // success
//	assert.NotEqual(t, "ABC", "CBA") // success
//	assert.NotEqual(t, 1, uint64(1)) // success
//	assert.NotEqual(t, 1, 1) // fail
//	assert.NotEqual(t, "ABC", "ABC") // fail
//	assert.NotEqual(t, 1, int64(1)) // fail
func NotEqual(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotEqual(t, false, actual, expect, message...)
}

// NotEqualNow tests the inequality between actual and expected parameters, and it'll stop the
// execution if they are equal.
//
//	assert.NotEqualNow(t, 1, 0) // success
//	assert.NotEqualNow(t, "ABC", "CBA") // success
//	assert.NotEqualNow(t, 1, 1) // fail and terminate
//	// never run
func NotEqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotEqual(t, true, actual, expect, message...)
}

// HasPrefixString tests whether the string has the prefix string or not, and it set the result to
// fail if the string does not have the prefix string.
//
//	assert.HasPrefixString(t, "Hello world", "") // success
//	assert.HasPrefixString(t, "Hello world", "Hello") // success
//	assert.HasPrefixString(t, "Hello world", "world") // fail
//	assert.HasPrefixString(t, "Hello world", "hello") // fail
func HasPrefixString(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryHasPrefixString(t, false, str, prefix, message...)
}

// HasPrefixStringNow tests whether the string has the prefix string or not, and it will terminate
// the execution if the string does not have the prefix string.
//
//	assert.HasPrefixStringNow(t, "Hello world", "") // success
//	assert.HasPrefixStringNow(t, "Hello world", "Hello") // success
//	assert.HasPrefixStringNow(t, "Hello world", "hello") // fail and stop the execution
//	// never runs
func HasPrefixStringNow(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryHasPrefixString(t, true, str, prefix, message...)
}

// NotHasPrefixString tests whether the string has the prefix string or not, and it set the result
// to fail if the string have the prefix string.
//
//	assert.NotHasPrefixString(t, "Hello world", "hello") // success
//	assert.NotHasPrefixString(t, "Hello world", "world") // success
//	assert.NotHasPrefixString(t, "Hello world", "") // fail
//	assert.NotHasPrefixString(t, "Hello world", "Hello") // fail
func NotHasPrefixString(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryNotHasPrefixString(t, false, str, prefix, message...)
}

// NotHasPrefixStringNow tests whether the string has the prefix string or not, and it will
// terminate the execution if the string have the prefix string.
//
//	assert.NotHasPrefixStringNow(t, "Hello world", "hello") // success
//	assert.NotHasPrefixStringNow(t, "Hello world", "world") // success
//	assert.NotHasPrefixStringNow(t, "Hello world", "Hello") // fail and stop the execution
//	// never runs
func NotHasPrefixStringNow(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryNotHasPrefixString(t, true, str, prefix, message...)
}

// HasSuffixString tests whether the string has the suffix string or not, and it set the result to
// fail if the string does not have the suffix string.
//
//	assert.HasSuffixString(t, "Hello world", "") // success
//	assert.HasSuffixString(t, "Hello world", "world") // success
//	assert.HasSuffixString(t, "Hello world", "World") // fail
//	assert.HasSuffixString(t, "Hello world", "hello") // fail
func HasSuffixString(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryHasSuffixString(t, false, str, suffix, message...)
}

// HasSuffixStringNow tests whether the string has the suffix string or not, and it will terminate
// the execution if the string does not have the suffix string.
//
//	assert.HasSuffixStringNow(t, "Hello world", "") // success
//	assert.HasSuffixStringNow(t, "Hello world", "world") // success
//	assert.HasSuffixStringNow(t, "Hello world", "World") // fail and stop the execution
//	// never runs
func HasSuffixStringNow(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryHasSuffixString(t, true, str, suffix, message...)
}

// NotHasSuffixString tests whether the string has the suffix string or not, and it set the result
// to fail if the string have the suffix string.
//
//	assert.NotHasSuffixString(t, "Hello world", "Hello") // success
//	assert.NotHasSuffixString(t, "Hello world", "World") // success
//	assert.NotHasSuffixString(t, "Hello world", "") // fail
//	assert.NotHasSuffixString(t, "Hello world", "world") // fail
func NotHasSuffixString(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryNotHasSuffixString(t, false, str, suffix, message...)
}

// NotHasSuffixStringNow tests whether the string has the suffix string or not, and it will
// terminate the execution if the string have the suffix string.
//
//	assert.NotHasSuffixStringNow(t, "Hello world", "hello") // success
//	assert.NotHasSuffixStringNow(t, "Hello world", "World") // success
//	assert.NotHasSuffixStringNow(t, "Hello world", "world") // fail and stop the execution
//	// never runs
func NotHasSuffixStringNow(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryNotHasSuffixString(t, true, str, suffix, message...)
}

// IsError tests whether the error matches the target or not. It'll set the result to fail if the
// error does not match to the target error, and it doesn't stop the execution.
//
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	err3 := errors.New("error 3")
//	assert.IsError(t, err1, err1) // success
//	assert.IsError(t, err1, err2) // fail
//	assert.IsError(t, errors.Join(err1, err2), err1) // success
//	assert.IsError(t, errors.Join(err1, err2), err2) // success
//	assert.IsError(t, errors.Join(err1, err2), err3) // fail
func IsError(t *testing.T, err, expected error, message ...any) error {
	return isError(t, false, err, expected, message...)
}

// IsErrorNow tests whether the error matches the target or not. It'll set the result to fail and
// stop the execution if the error does not match to the target error.
//
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	assert.IsErrorNow(t, errors.Join(err1, err2), err1) // success
//	assert.IsErrorNow(t, errors.Join(err1, err2), err2) // success
//	assert.IsErrorNow(t, err1, err1) // success
//	assert.IsErrorNow(t, err1, err2) // fail
//	// never runs
func IsErrorNow(t *testing.T, err, expected error, message ...any) error {
	return isError(t, true, err, expected, message...)
}

// NotIsError tests whether the error matches the target or not. It'll set the result to fail if
// the error matches to the target error, and it doesn't stop the execution.
//
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	err3 := errors.New("error 3")
//	assert.NotIsError(t, err1, err2) // success
//	assert.NotIsError(t, err1, err1) // fail
//	assert.NotIsError(t, errors.Join(err1, err2), err3) // success
//	assert.NotIsError(t, errors.Join(err1, err2), err1) // fail
//	assert.NotIsError(t, errors.Join(err1, err2), err2) // fail
func NotIsError(t *testing.T, err, unexpected error, message ...any) error {
	return notIsError(t, false, err, unexpected, message...)
}

// NotIsErrorNow tests whether the error matches the target or not. It'll set the result to fail
// and stop the execution if the error matches to the target error.
//
//	err1 := errors.New("error 1")
//	err2 := errors.New("error 2")
//	err3 := errors.new("error 3")
//	assert.NotIsErrorNow(t, errors.Join(err1, err2), err3) // success
//	assert.NotIsErrorNow(t, err1, err2) // fail
//	assert.NotIsErrorNow(t, err1, err1) // fail and terminate
//	// never runs
func NotIsErrorNow(t *testing.T, err, unexpected error, message ...any) error {
	return notIsError(t, true, err, unexpected, message...)
}

// MapHasKey tests whether the map contains the specified key or not, it will fail if the map does
// not contain the key, or the type of the key cannot assign to the type of the key of the map.
//
//	assert.MapHasKey(t, map[string]int{"a":1}, "a") // success
//	assert.MapHasKey(t, map[string]int{"a":1}, "b") // fail
//	assert.MapHasKey(t, map[string]int{"a":1}, 1) // fail
func MapHasKey(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryMapHasKey(t, false, m, key, message...)
}

// MapHasKeyNow tests whether the map contains the specified key or not, and it will terminate the
// execution if the test fails. It will fail if the map does not contain the key, or the type of
// the key cannot assign to the type of the key of the map.
//
//	assert.MapHasKeyNow(t, map[string]int{"a":1}, "a") // success
//	assert.MapHasKeyNow(t, map[string]int{"a":1}, "b") // fail and terminate
//	// never run
func MapHasKeyNow(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryMapHasKey(t, true, m, key, message...)
}

// NotMapHasKey tests whether the map contains the specified key or not, it will fail if the map
// contain the key. It will also set the test result to success if the type of the key cannot
// assign to the type of the key of the map.
//
//	assert.NotMapHasKey(t, map[string]int{"a":1}, "b") // success
//	assert.NotMapHasKey(t, map[string]int{"a":1}, 1) // success
//	assert.NotMapHasKey(t, map[string]int{"a":1}, "a") // fail
func NotMapHasKey(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryNotMapHasKey(t, false, m, key, message...)
}

// NotMapHasKeyNow tests whether the map contains the specified key or not, it will fail if the map
// contain the key, and it will terminate the execution if the test fails. It will also set the
// test result to success if the type of the key cannot assign to the type of the key of the map.
//
//	assert.NotMapHasKeyNow(t, map[string]int{"a":1}, "b") // success
//	assert.NotMapHasKeyNow(t, map[string]int{"a":1}, 1) // success
//	assert.NotMapHasKeyNow(t, map[string]int{"a":1}, "a") // fail and terminate
//	// never run
func NotMapHasKeyNow(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryNotMapHasKey(t, true, m, key, message...)
}

// MapHasValue tests whether the map contains the specified value or not, it will fail if the map
// does not contain the value, or the type of the value cannot assign to the type of the values of
// the map.
//
//	assert.MapHasValue(t, map[string]int{"a":1}, 1) // success
//	assert.MapHasValue(t, map[string]int{"a":1}, 2) // fail
//	assert.MapHasValue(t, map[string]int{"a":1}, "a") // fail
func MapHasValue(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryMapHasValue(t, false, m, value, message...)
}

// MapHasValueNow tests whether the map contains the specified value or not, and it will terminate
// the execution if the test fails. It will fail if the map does not contain the value, or the type
// of the value cannot assign to the type of the value of the map.
//
//	assert.MapHasValueNow(t, map[string]int{"a":1}, 1) // success
//	assert.MapHasValueNow(t, map[string]int{"a":1}, 2) // fail and terminate
//	// never run
func MapHasValueNow(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryMapHasValue(t, true, m, value, message...)
}

// NotMapHasValue tests whether the map contains the specified value or not, it will fail if the
// map contain the value. It will also set the test result to success if the type of the value
// cannot assign to the type of the value of the map.
//
//	assert.NotMapHasValue(t, map[string]int{"a":1}, 2) // success
//	assert.NotMapHasValue(t, map[string]int{"a":1}, "a") // success
//	assert.NotMapHasValue(t, map[string]int{"a":1}, 1) // fail
func NotMapHasValue(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryNotMapHasValue(t, false, m, value, message...)
}

// NotMapHasValueNow tests whether the map contains the specified value or not, it will fail if the
// map contain the value, and it will terminate the execution if the test fails. It will also set
// the test result to success if the type of the value cannot assign to the type of the value of
// the map.
//
//	assert.NotMapHasValueNow(t, map[string]int{"a":1}, 2) // success
//	assert.NotMapHasValueNow(t, map[string]int{"a":1}, "a") // success
//	assert.NotMapHasValueNow(t, map[string]int{"a":1}, 1) // fail and terminate
//	// never run
func NotMapHasValueNow(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryNotMapHasValue(t, true, m, value, message...)
}

// Match tests whether the string matches the regular expression or not.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assert.Match(t, "http://example.com", pattern) // success
//	assert.Match(t, "example.com", pattern) // fail
func Match(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, false, val, pattern, "", message...)
}

// MatchNow tests whether the string matches the regular expression or not, and it will terminate
// the execution if it does not match.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assert.MatchNow(t, "http://example.com", pattern) // success
//	assert.MatchNow(t, "example.com", pattern) // fail and terminate
//	// never run
func MatchNow(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not. It will panic if the pattern is not a valid regular expression.
//
//	assert.MatchString(t, "http://example.com", `^https?:\/\/`) // success
//	assert.MatchString(t, "example.com", `^https?:\/\/`) // fail
func MatchString(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, false, val, nil, pattern, message...)
}

// MatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not. It will terminate the execution if it does not match, and it will panic if
// the pattern is not a valid regular expression.
//
//	assert.MatchStringNow(t, "http://example.com", `^https?:\/\/`) // success
//	assert.MatchStringNow(t, "example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func MatchStringNow(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, true, val, nil, pattern, message...)
}

// NotMatch tests whether the string matches the regular expression or not, and it set the result
// to fail if the string matches the pattern.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assert.NotMatch(t, "example.com", pattern) // success
//	assert.NotMatch(t, "http://example.com", pattern) // fail
func NotMatch(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, false, val, pattern, "", message...)
}

// NotMatchNow tests whether the string matches the regular expression or not, and it will
// terminate the execution if the string matches the pattern.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assert.NotMatchNow(t, "example.com", pattern) // success
//	assert.NotMatchNow(t, "http://example.com", pattern) // fail and terminate
//	// never run
func NotMatchNow(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not, and it set the result to fail if the string matches the pattern. It will also panic if
// the pattern is not a valid regular expression.
//
//	assert.NotMatchString(t, "example.com", `^https?:\/\/`) // success
//	assert.NotMatchString(t, "http://example.com", `^https?:\/\/`) // fail
func NotMatchString(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, false, val, nil, pattern, message...)
}

// NotMatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not, and it set the result to fail if the string matches the pattern. It will
// terminate the execution if the string matches the pattern, and it will panic if the pattern is
// not a valid regular expression.
//
//	assert.NotMatchStringNow(t, "example.com", `^https?:\/\/`) // success
//	assert.NotMatchStringNow(t, "http://example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func NotMatchStringNow(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, true, val, nil, pattern, message...)
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
//	var err error // nil
//	assert.Nil(t, err) // success
//
//	err = errors.New("some error")
//	assert.Nil(t, err) // fail
func Nil(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNil(t, false, val, message...)
}

// NilNow tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is not nil.
//
//	var err error // nil
//	assert.NilNow(t, err) // success
//
//	err = errors.New("some error")
//	assert.NilNow(t, err) // fail and terminate
//	// never run
func NilNow(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNil(t, true, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
//	var err error // nil
//	assert.NotNil(t, err) // fail
//
//	err = errors.New("some error")
//	assert.NotNil(t, err) // success
func NotNil(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNotNil(t, false, val, message...)
}

// NotNilNow tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
//
// This function will set the result to fail, and stop the execution if the value is nil.
//
//	var err error = errors.New("some error")
//	assert.NotNilNow(t, err) // success
//
//	err = nil
//	assert.NotNilNow(t, err) // fail and terminate
//	// never run
func NotNilNow(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNotNil(t, true, val, message...)
}

// Panic expects the function fn to panic, and it'll set the result to fail if the function doesn't
// panic.
//
//	assert.Panic(t, func() {
//	  panic("some error")
//	}) // success
//
//	assert.Panic(t, func() {
//	  // no panic
//	}) // fail
func Panic(t *testing.T, fn func(), message ...any) error {
	t.Helper()

	return tryPanic(t, false, fn, message...)
}

// PanicNow expects the function fn to panic. It'll set the result to fail if the function doesn't
// panic, and stop the execution.
//
//	assert.PanicNow(t, func() {
//	  panic("some error")
//	}) // success
//
//	assert.PanicNow(t, func() {
//	  // no panic
//	}) // fail
//	// never run
func PanicNow(t *testing.T, fn func(), message ...any) error {
	t.Helper()

	return tryPanic(t, true, fn, message...)
}

// NotPanic asserts that the function fn does not panic, and it'll set the result to fail if the
// function panic.
//
//	assert.NotPanic(t, func() {
//	  // no panic
//	}) // success
//
//	assert.NotPanic(t, func() {
//	  panic("some error")
//	}) // fail
func NotPanic(t *testing.T, fn func(), message ...any) error {
	t.Helper()

	return tryNotPanic(t, false, fn, message...)
}

// NotPanicNow asserts that the function fn does not panic. It'll set the result to fail if the
// function panic, and it also stops the execution.
//
//	assert.NotPanicNow(t, func() {
//	  // no panic
//	}) // success
//
//	assert.NotPanicNow(t, func() {
//	  panic("some error")
//	}) // fail and terminate
//	// never run
func NotPanicNow(t *testing.T, fn func(), message ...any) error {
	t.Helper()

	return tryNotPanic(t, true, fn, message...)
}

// PanicOf expects the function fn to panic by the expected error. If the function does not panic
// or panic for another reason, it will set the result to fail.
//
//	assert.PanicOf(t, func() {
//	  panic("expected error")
//	}, "expected error") // success
//	assert.PanicOf(t, func() {
//	  panic("unexpected error")
//	}, "expected error") // fail
//	assert.PanicOf(t, func() {
//	  // ..., no panic
//	}, "expected error") // fail
func PanicOf(t *testing.T, fn func(), expectErr any, message ...any) error {
	t.Helper()

	return tryPanicOf(t, false, fn, expectErr, message...)
}

// PanicOfNow expects the function fn to panic by the expected error. If the function does not
// panic or panic for another reason, it will set the result to fail and terminate the execution.
//
//	assert.PanicOfNow(t, func() {
//	  panic("expected error")
//	}, "expected error") // success
//	assert.PanicOfNow(t, func() {
//	  panic("unexpected error")
//	}, "expected error") // fail and terminated
//	// never runs
func PanicOfNow(t *testing.T, fn func(), expectErr any, message ...any) error {
	t.Helper()

	return tryPanicOf(t, true, fn, expectErr, message...)
}

// NotPanicOf expects the function fn not panic, or the function does not panic by the unexpected
// error. If the function panics by the unexpected error, it will set the result to fail.
//
//	assert.NotPanicOf(t, func() {
//	  panic("other error")
//	}, "unexpected error") // success
//	assert.NotPanicOf(t, func() {
//	  // ..., no panic
//	}, "unexpected error") // success
//	assert.NotPanicOf(t, func() {
//	  panic("unexpected error")
//	}, "unexpected error") // fail
func NotPanicOf(t *testing.T, fn func(), unexpectedErr any, message ...any) error {
	t.Helper()

	return tryNotPanicOf(t, false, fn, unexpectedErr, message...)
}

// NotPanicOfNow expects the function fn not panic, or the function does not panic by the
// unexpected error. If the function panics by the unexpected error, it will set the result to fail
// and stop the execution.
//
//	assert.NotPanicOfNow(t, func() {
//	  panic("other error")
//	}, "unexpected error") // success
//	assert.NotPanicOfNow(t, func() {
//	  // ..., no panic
//	}, "unexpected error") // success
//	assert.NotPanicOfNow(t, func() {
//	  panic("unexpected error")
//	}, "unexpected error") // fail and terminate
//	// never runs
func NotPanicOfNow(t *testing.T, fn func(), unexpectedErr any, message ...any) error {
	t.Helper()

	return tryNotPanicOf(t, true, fn, unexpectedErr, message...)
}

// True tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
//	assert.True(t, 1) // success
//	assert.True(t, "test") // success
//	assert.True(t, 0) // fail
//	assert.True(t, "") // fail
func True(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryTrue(t, false, val, message...)
}

// TrueNow tests whether a value is truthy or not. It'll set the result to fail if the value is a
// false value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
// The function will stop the execution if the value is falsy.
//
//	assert.TrueNow(t, 1) // success
//	assert.TrueNow(t, "test") // success
//	assert.TrueNow(t, "") // fail and terminate
//	// never run
func TrueNow(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryTrue(t, true, val, message...)
}

// NotTrue tests whether a value is truthy or not. It'll set the result to fail if the value is a
// truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
//	assert.NotTrue(t, 0) // success
//	assert.NotTrue(t, "") // success
//	assert.NotTrue(t, 1) // fail
//	assert.NotTrue(t, "test") // fail
func NotTrue(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNotTrue(t, false, val, message...)
}

// NotTrueNow tests whether a value is truthy or not. It'll set the result to fail if the value is
// a truthy value. For most types of value, a falsy value is the zero value for its type. For a
// slice, a truthy value should not be nil, and its length must be greater than 0. For nil, the
// value is always falsy.
//
// The function will stop the execution if the value is truthy.
//
//	assert.NotTrueNow(t, 0) // success
//	assert.NotTrueNow(t, "") // success
//	assert.NotTrueNow(t, "test") // fail and terminate
//	// never run
func NotTrueNow(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNotTrue(t, true, val, message...)
}
