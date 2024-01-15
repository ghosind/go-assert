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
//	assertion.ContainsElement([]int{1, 2, 3}, 1) // success
//	assertion.ContainsElement([]int{1, 2, 3}, 3) // success
//	assertion.ContainsElement([]int{1, 2, 3}, 4) // fail
func ContainsElement(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryContainsElement(t, false, source, expect, message...)
}

// ContainsElementNow tests whether the array or slice contains the specified element or not, and
// it will terminate the execution if the array or slice does not contain the specified element.
// It'll panic if the `source` is not an array or a slice.
//
//	assertion.ContainsElementNow([]int{1, 2, 3}, 1) // success
//	assertion.ContainsElementNow([]int{1, 2, 3}, 3) // success
//	assertion.ContainsElementNow([]int{1, 2, 3}, 4) // fail and stop the execution
//	// never runs
func ContainsElementNow(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryContainsElement(t, true, source, expect, message...)
}

// NotContainsElement tests whether the array or slice contains the specified element or not, and
// it set the result to fail if the array or slice contains the specified element. It'll panic if
// the `source` is not an array or a slice.
//
//	assertion.NotContainsElement([]int{1, 2, 3}, 4) // success
//	assertion.NotContainsElement([]int{1, 2, 3}, 0) // success
//	assertion.NotContainsElement([]int{1, 2, 3}, 1) // fail
func NotContainsElement(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryNotContainsElement(t, false, source, expect, message...)
}

// NotContainsElementNow tests whether the array or slice contains the specified element or not,
// and it will terminate the execution if the array or slice contains the specified element. It'll
// panic if the `source` is not an array or a slice.
//
//	assertion.NotContainsElementNow([]int{1, 2, 3}, 4) // success
//	assertion.NotContainsElementNow([]int{1, 2, 3}, 0) // success
//	assertion.NotContainsElementNow([]int{1, 2, 3}, 1) // fail and stop the execution
//	// never runs
func NotContainsElementNow(t *testing.T, source, expect any, message ...any) error {
	t.Helper()

	return tryNotContainsElement(t, true, source, expect, message...)
}

// ContainsString tests whether the string contains the substring or not, and it set the result to
// fail if the string does not contains the substring.
//
//	ContainsString(t, "Hello world", "") // success
//	ContainsString(t, "Hello world", "Hello") // success
//	ContainsString(t, "Hello world", "world") // success
//	ContainsString(t, "Hello world", "hello") // fail
func ContainsString(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryContainsString(t, false, str, substr, message...)
}

// ContainsStringNow tests whether the string contains the substring or not, and it will terminate the
// execution if the string does not contains the substring.
//
//	ContainsStringNow(t, "Hello world", "") // success
//	ContainsStringNow(t, "Hello world", "Hello") // success
//	ContainsStringNow(t, "Hello world", "world") // success
//	ContainsStringNow(t, "Hello world", "hello") // fail and stop the execution
//	// never runs
func ContainsStringNow(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryContainsString(t, true, str, substr, message...)
}

// Gt compares the values and sets the result to false if the first value is not greater than to
// the second value.
//
//	Gt(t, 2, 1) // success
//	Gt(t, 3.14, 1.68) // success
//	Gt(t, "BCD", "ABC") // success
//	Gt(t, 2, 2) // fail
//	Gt(t, 1, 2) // fail
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
//	GtNow(t, 2, 1) // success
//	GtNow(t, 3.14, 1.68) // success
//	GtNow(t, "BCD", "ABC") // success
//	GtNow(t, 1, 2) // fail and terminate
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
//	Gte(t, 2, 1) // success
//	Gte(t, 3.14, 1.68) // success
//	Gte(t, "BCD", "ABC") // success
//	Gte(t, 2, 2) // success
//	Gte(t, 1, 2) // fail
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
//	GteNow(t, 2, 1) // success
//	GteNow(t, 3.14, 1.68) // success
//	GteNow(t, "BCD", "ABC") // success
//	GteNow(t, 2, 2) // success
//	GteNow(t, 1, 2) // fail and terminate
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
//	Lt(t, 1, 2) // success
//	Lt(t, 1.68, 3.14) // success
//	Lt(t, "ABC", "BCD") // success
//	Lt(t, 2, 2) // fail
//	Lt(t, 2, 1) // fail
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
//	LtNow(t, 1, 2) // success
//	LtNow(t, 1.68, 3.14) // success
//	LtNow(t, "ABC", "BCD") // success
//	LtNow(t, 2, 1) // fail and terminate
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
//	Lte(t, 1, 2) // success
//	Lte(t, 1.68, 3.14) // success
//	Lte(t, "ABC", "BCD") // success
//	Lte(t, 2, 2) // success
//	Lte(t, 2, 1) // fail
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
//	LteNow(t, 1, 2) // success
//	LteNow(t, 1.68, 3.14) // success
//	LteNow(t, "ABC", "BCD") // success
//	LteNow(t, 2, 2) // success
//	LteNow(t, 2, 1) // fail and terminate
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
//	NotContainsString(t, "Hello world", "") // fail
//	NotContainsString(t, "Hello world", "Hello") // fail
//	NotContainsString(t, "Hello world", "world") // fail
//	NotContainsString(t, "Hello world", "hello") // success
func NotContainsString(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryNotContainsString(t, false, str, substr, message...)
}

// NotContainsStringNow tests whether the string contains the substring or not, and it will terminate the
// execution if the string contains the substring.
//
//	NotContainsStringNow(t, "Hello world", "hello") // success
//	NotContainsStringNow(t, "Hello world", "Hello") // fail and stop the execution
//	// never runs
func NotContainsStringNow(t *testing.T, str, substr string, message ...any) error {
	t.Helper()

	return tryNotContainsString(t, true, str, substr, message...)
}

// DeepEqual tests the deep equality between actual and expect parameters. It'll set the result to
// fail if they are not deeply equal, and it doesn't stop the execution.
//
//	DeepEqual(t, 1, 1) // success
//	DeepEqual(t, "ABC", "ABC") // success
//	DeepEqual(t, 1, 0) // fail
//	DeepEqual(t, 1, int64(1)) // fail
func DeepEqual(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryDeepEqual(t, false, actual, expect, message...)
}

// DeepEqualNow tests the deep equality between actual and expect parameters, and it'll stop the
// execution if they are not deeply equal.
//
//	DeepEqualNow(t, 1, 1) // success
//	DeepEqualNow(t, "ABC", "ABC") // success
//	DeepEqualNow(t, 1, int64(1)) // fail and terminate
//	// never run
func DeepEqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryDeepEqual(t, true, actual, expect, message...)
}

// NotDeepEqual tests the deep inequality between actual and expected parameters. It'll set the
// result to fail if they are deeply equal, but it doesn't stop the execution.
//
//	NotDeepEqual(t, 1, 0) // success
//	NotDeepEqual(t, 1, int64(1)) // success
//	NotDeepEqual(t, 1, 1) // fail
//	NotDeepEqual(t, "ABC", "ABC") // fail
func NotDeepEqual(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotDeepEqual(t, false, actual, expect, message...)
}

// NotDeepEqualNow tests the deep inequality between actual and expected parameters, and it'll stop
// the execution if they are deeply equal.
//
//	NotDeepEqual(t, 1, 0) // success
//	NotDeepEqual(t, 1, int64(1)) // success
//	NotDeepEqual(t, "ABC", "ABC") // fail and terminate
//	// never run
func NotDeepEqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotDeepEqual(t, true, actual, expect, message...)
}

// Equal tests the equality between actual and expect parameters. It'll set the result to fail if
// they are not equal, and it doesn't stop the execution.
//
//	Equal(t, 1, 1) // success
//	Equal(t, "ABC", "ABC") // success
//	Equal(t, 1, int64(1)) // success
//	Equal(t, 1, uint64(1)) // fail
//	Equal(t, 1, 0) // fail
func Equal(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryEqual(t, false, actual, expect, message...)
}

// EqualNow tests the equality between actual and expect parameters, and it'll stop the execution
// if they are not equal.
//
//	EqualNow(t, 1, 1) // success
//	EqualNow(t, "ABC", "ABC") // success
//	EqualNow(t, 1, int64(1)) // success
//	EqualNow(t, 1, 0) // fail and terminate
//	never run
func EqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryEqual(t, true, actual, expect, message...)
}

// NotEqual tests the inequality between actual and expected parameters. It'll set the result to
// fail if they are equal, but it doesn't stop the execution.
//
//	NotEqual(t, 1, 0) // success
//	NotEqual(t, "ABC", "CBA") // success
//	NotEqual(t, 1, uint64(1)) // success
//	NotEqual(t, 1, 1) // fail
//	NotEqual(t, "ABC", "ABC") // fail
//	NotEqual(t, 1, int64(1)) // fail
func NotEqual(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotEqual(t, false, actual, expect, message...)
}

// NotEqualNow tests the inequality between actual and expected parameters, and it'll stop the
// execution if they are equal.
//
//	NotEqualNow(t, 1, 0) // success
//	NotEqualNow(t, "ABC", "CBA") // success
//	NotEqualNow(t, 1, 1) // fail and terminate
//	// never run
func NotEqualNow(t *testing.T, actual, expect any, message ...any) error {
	t.Helper()

	return tryNotEqual(t, true, actual, expect, message...)
}

// HasPrefixString tests whether the string has the prefix string or not, and it set the result to
// fail if the string does not have the prefix string.
//
//	HasPrefixString(t, "Hello world", "") // success
//	HasPrefixString(t, "Hello world", "Hello") // success
//	HasPrefixString(t, "Hello world", "world") // fail
//	HasPrefixString(t, "Hello world", "hello") // fail
func HasPrefixString(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryHasPrefixString(t, false, str, prefix, message...)
}

// HasPrefixStringNow tests whether the string has the prefix string or not, and it will terminate
// the execution if the string does not have the prefix string.
//
//	HasPrefixStringNow(t, "Hello world", "") // success
//	HasPrefixStringNow(t, "Hello world", "Hello") // success
//	HasPrefixStringNow(t, "Hello world", "hello") // fail and stop the execution
//	// never runs
func HasPrefixStringNow(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryHasPrefixString(t, true, str, prefix, message...)
}

// NotHasPrefixString tests whether the string has the prefix string or not, and it set the result
// to fail if the string have the prefix string.
//
//	NotHasPrefixString(t, "Hello world", "hello") // success
//	NotHasPrefixString(t, "Hello world", "world") // success
//	NotHasPrefixString(t, "Hello world", "") // fail
//	NotHasPrefixString(t, "Hello world", "Hello") // fail
func NotHasPrefixString(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryNotHasPrefixString(t, false, str, prefix, message...)
}

// NotHasPrefixStringNow tests whether the string has the prefix string or not, and it will
// terminate the execution if the string have the prefix string.
//
//	NotHasPrefixStringNow(t, "Hello world", "hello") // success
//	NotHasPrefixStringNow(t, "Hello world", "world") // success
//	NotHasPrefixStringNow(t, "Hello world", "Hello") // fail and stop the execution
//	// never runs
func NotHasPrefixStringNow(t *testing.T, str, prefix string, message ...any) error {
	t.Helper()

	return tryNotHasPrefixString(t, true, str, prefix, message...)
}

// HasSuffixString tests whether the string has the suffix string or not, and it set the result to
// fail if the string does not have the suffix string.
//
//	HasSuffixString(t, "Hello world", "") // success
//	HasSuffixString(t, "Hello world", "world") // success
//	HasSuffixString(t, "Hello world", "World") // fail
//	HasSuffixString(t, "Hello world", "hello") // fail
func HasSuffixString(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryHasSuffixString(t, false, str, suffix, message...)
}

// HasSuffixStringNow tests whether the string has the suffix string or not, and it will terminate
// the execution if the string does not have the suffix string.
//
//	HasSuffixStringNow(t, "Hello world", "") // success
//	HasSuffixStringNow(t, "Hello world", "world") // success
//	HasSuffixStringNow(t, "Hello world", "World") // fail and stop the execution
//	// never runs
func HasSuffixStringNow(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryHasSuffixString(t, true, str, suffix, message...)
}

// NotHasSuffixString tests whether the string has the suffix string or not, and it set the result
// to fail if the string have the suffix string.
//
//	NotHasSuffixString(t, "Hello world", "Hello") // success
//	NotHasSuffixString(t, "Hello world", "World") // success
//	NotHasSuffixString(t, "Hello world", "") // fail
//	NotHasSuffixString(t, "Hello world", "world") // fail
func NotHasSuffixString(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryNotHasSuffixString(t, false, str, suffix, message...)
}

// NotHasSuffixStringNow tests whether the string has the suffix string or not, and it will
// terminate the execution if the string have the suffix string.
//
//	NotHasSuffixStringNow(t, "Hello world", "hello") // success
//	NotHasSuffixStringNow(t, "Hello world", "World") // success
//	NotHasSuffixStringNow(t, "Hello world", "world") // fail and stop the execution
//	// never runs
func NotHasSuffixStringNow(t *testing.T, str, suffix string, message ...any) error {
	t.Helper()

	return tryNotHasSuffixString(t, true, str, suffix, message...)
}

// MapHasKey tests whether the map contains the specified key or not, it will fail if the map does
// not contain the key, or the type of the key cannot assign to the type of the key of the map.
//
//	assertion.MapHasKey(map[string]int{"a":1}, "a") // success
//	assertion.MapHasKey(map[string]int{"a":1}, "b") // fail
//	assertion.MapHasKey(map[string]int{"a":1}, 1) // fail
func MapHasKey(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryMapHasKey(t, false, m, key, message...)
}

// MapHasKeyNow tests whether the map contains the specified key or not, and it will terminate the
// execution if the test fails. It will fail if the map does not contain the key, or the type of
// the key cannot assign to the type of the key of the map.
//
//	assertion.MapHasKeyNow(map[string]int{"a":1}, "a") // success
//	assertion.MapHasKeyNow(map[string]int{"a":1}, "b") // fail and terminate
//	// never run
func MapHasKeyNow(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryMapHasKey(t, true, m, key, message...)
}

// NotMapHasKey tests whether the map contains the specified key or not, it will fail if the map
// contain the key. It will also set the test result to success if the type of the key cannot
// assign to the type of the key of the map.
//
//	assertion.NotMapHasKey(map[string]int{"a":1}, "b") // success
//	assertion.NotMapHasKey(map[string]int{"a":1}, 1) // success
//	assertion.NotMapHasKey(map[string]int{"a":1}, "a") // fail
func NotMapHasKey(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryNotMapHasKey(t, false, m, key, message...)
}

// NotMapHasKeyNow tests whether the map contains the specified key or not, it will fail if the map
// contain the key, and it will terminate the execution if the test fails. It will also set the
// test result to success if the type of the key cannot assign to the type of the key of the map.
//
//	assertion.NotMapHasKeyNow(map[string]int{"a":1}, "b") // success
//	assertion.NotMapHasKeyNow(map[string]int{"a":1}, 1) // success
//	assertion.NotMapHasKeyNow(map[string]int{"a":1}, "a") // fail and terminate
//	// never run
func NotMapHasKeyNow(t *testing.T, m, key any, message ...any) error {
	t.Helper()

	return tryNotMapHasKey(t, true, m, key, message...)
}

// MapHasValue tests whether the map contains the specified value or not, it will fail if the map
// does not contain the value, or the type of the value cannot assign to the type of the values of
// the map.
//
//	assertion.MapHasValue(map[string]int{"a":1}, 1) // success
//	assertion.MapHasValue(map[string]int{"a":1}, 2) // fail
//	assertion.MapHasValue(map[string]int{"a":1}, "a") // fail
func MapHasValue(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryMapHasValue(t, false, m, value, message...)
}

// MapHasValueNow tests whether the map contains the specified value or not, and it will terminate
// the execution if the test fails. It will fail if the map does not contain the value, or the type
// of the value cannot assign to the type of the value of the map.
//
//	assertion.MapHasValueNow(map[string]int{"a":1}, 1) // success
//	assertion.MapHasValueNow(map[string]int{"a":1}, 2) // fail and terminate
//	// never run
func MapHasValueNow(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryMapHasValue(t, true, m, value, message...)
}

// NotMapHasValue tests whether the map contains the specified value or not, it will fail if the
// map contain the value. It will also set the test result to success if the type of the value
// cannot assign to the type of the value of the map.
//
//	assertion.NotMapHasValue(map[string]int{"a":1}, 2) // success
//	assertion.NotMapHasValue(map[string]int{"a":1}, "a") // success
//	assertion.NotMapHasValue(map[string]int{"a":1}, 1) // fail
func NotMapHasValue(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryNotMapHasValue(t, false, m, value, message...)
}

// NotMapHasValueNow tests whether the map contains the specified value or not, it will fail if the
// map contain the value, and it will terminate the execution if the test fails. It will also set
// the test result to success if the type of the value cannot assign to the type of the value of
// the map.
//
//	assertion.NotMapHasValueNow(map[string]int{"a":1}, 2) // success
//	assertion.NotMapHasValueNow(map[string]int{"a":1}, "a") // success
//	assertion.NotMapHasValueNow(map[string]int{"a":1}, 1) // fail and terminate
//	// never run
func NotMapHasValueNow(t *testing.T, m, value any, message ...any) error {
	t.Helper()

	return tryNotMapHasValue(t, true, m, value, message...)
}

// Match tests whether the string matches the regular expression or not.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	Match(t, "http://example.com", pattern) // success
//	Match(t, "example.com", pattern) // fail
func Match(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, false, val, pattern, "", message...)
}

// MatchNow tests whether the string matches the regular expression or not, and it will terminate
// the execution if it does not match.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	MatchNow(t, "http://example.com", pattern) // success
//	MatchNow(t, "example.com", pattern) // fail and terminate
//	// never run
func MatchNow(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not. It will panic if the pattern is not a valid regular expression.
//
//	MatchString(t, "http://example.com", `^https?:\/\/`) // success
//	MatchString(t, "example.com", `^https?:\/\/`) // fail
func MatchString(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, false, val, nil, pattern, message...)
}

// MatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not. It will terminate the execution if it does not match, and it will panic if
// the pattern is not a valid regular expression.
//
//	MatchStringNow(t, "http://example.com", `^https?:\/\/`) // success
//	MatchStringNow(t, "example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func MatchStringNow(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryMatchRegexp(t, true, val, nil, pattern, message...)
}

// NotMatch tests whether the string matches the regular expression or not, and it set the result
// to fail if the string matches the pattern.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	NotMatch(t, "example.com", pattern) // success
//	NotMatch(t, "http://example.com", pattern) // fail
func NotMatch(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, false, val, pattern, "", message...)
}

// NotMatchNow tests whether the string matches the regular expression or not, and it will
// terminate the execution if the string matches the pattern.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	NotMatchNow(t, "example.com", pattern) // success
//	NotMatchNow(t, "http://example.com", pattern) // fail and terminate
//	// never run
func NotMatchNow(t *testing.T, val string, pattern *regexp.Regexp, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not, and it set the result to fail if the string matches the pattern. It will also panic if
// the pattern is not a valid regular expression.
//
//	NotMatchString(t, "example.com", `^https?:\/\/`) // success
//	NotMatchString(t, "http://example.com", `^https?:\/\/`) // fail
func NotMatchString(t *testing.T, val, pattern string, message ...any) error {
	t.Helper()

	return tryNotMatchRegexp(t, false, val, nil, pattern, message...)
}

// NotMatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not, and it set the result to fail if the string matches the pattern. It will
// terminate the execution if the string matches the pattern, and it will panic if the pattern is
// not a valid regular expression.
//
//	NotMatchStringNow(t, "example.com", `^https?:\/\/`) // success
//	NotMatchStringNow(t, "http://example.com", `^https?:\/\/`) // fail and terminate
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
//	Nil(t, err) // success
//
//	err = errors.New("some error")
//	Nil(t, err) // fail
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
//	NilNow(t, err) // success
//
//	err = errors.New("some error")
//	NilNow(t, err) // fail and terminate
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
//	NotNil(t, err) // fail
//
//	err = errors.New("some error")
//	NotNil(t, err) // success
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
//	NotNilNow(t, err) // success
//
//	err = nil
//	NotNilNow(t, err) // fail and terminate
//	// never run
func NotNilNow(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNotNil(t, true, val, message...)
}

// Panic expects the function fn to panic, and it'll set the result to fail if the function doesn't
// panic.
//
//	Panic(t, func() {
//	  panic("some error")
//	}) // success
//
//	Panic(t, func() {
//	  // no panic
//	}) // fail
func Panic(t *testing.T, fn func(), message ...any) error {
	t.Helper()

	return tryPanic(t, false, fn, message...)
}

// PanicNow expects the function fn to panic. It'll set the result to fail if the function doesn't
// panic, and stop the execution.
//
//	PanicNow(t, func() {
//	  panic("some error")
//	}) // success
//
//	PanicNow(t, func() {
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
//	NotPanic(t, func() {
//	  // no panic
//	}) // success
//
//	NotPanic(t, func() {
//	  panic("some error")
//	}) // fail
func NotPanic(t *testing.T, fn func(), message ...any) error {
	t.Helper()

	return tryNotPanic(t, false, fn, message...)
}

// NotPanicNow asserts that the function fn does not panic. It'll set the result to fail if the
// function panic, and it also stops the execution.
//
//	NotPanicNow(t, func() {
//	  // no panic
//	}) // success
//
//	NotPanicNow(t, func() {
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
//	PanicOf(t, func() {
//	  panic("expected error")
//	}, "expected error") // success
//	PanicOf(t, func() {
//	  panic("unexpected error")
//	}, "expected error") // fail
//	PanicOf(t, func() {
//	  // ..., no panic
//	}, "expected error") // fail
func PanicOf(t *testing.T, fn func(), expectErr any, message ...any) error {
	t.Helper()

	return tryPanicOf(t, false, fn, expectErr, message...)
}

// PanicOfNow expects the function fn to panic by the expected error. If the function does not
// panic or panic for another reason, it will set the result to fail and terminate the execution.
//
//	PanicOfNow(t, func() {
//	  panic("expected error")
//	}, "expected error") // success
//	PanicOfNow(t, func() {
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
//	NotPanicOf(t, func() {
//	  panic("other error")
//	}, "unexpected error") // success
//	NotPanicOf(t, func() {
//	  // ..., no panic
//	}, "unexpected error") // success
//	NotPanicOf(t, func() {
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
//	NotPanicOfNow(t, func() {
//	  panic("other error")
//	}, "unexpected error") // success
//	NotPanicOfNow(t, func() {
//	  // ..., no panic
//	}, "unexpected error") // success
//	NotPanicOfNow(t, func() {
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
//	True(t, 1) // success
//	True(t, "test") // success
//	True(t, 0) // fail
//	True(t, "") // fail
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
//	TrueNow(t, 1) // success
//	TrueNow(t, "test") // success
//	TrueNow(t, "") // fail and terminate
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
//	NotTrue(t, 0) // success
//	NotTrue(t, "") // success
//	NotTrue(t, 1) // fail
//	NotTrue(t, "test") // fail
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
//	NotTrueNow(t, 0) // success
//	NotTrueNow(t, "") // success
//	NotTrueNow(t, "test") // fail and terminate
//	// never run
func NotTrueNow(t *testing.T, val any, message ...any) error {
	t.Helper()

	return tryNotTrue(t, true, val, message...)
}
