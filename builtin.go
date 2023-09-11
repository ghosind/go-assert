package assert

import (
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
//	Equal(t, 1, uint64(1)) // success
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
//	EqualNow(t, 1, uint64(1)) // success
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
//	NotEqual(t, 1, 1) // fail
//	NotEqual(t, "ABC", "ABC") // fail
//	NotEqual(t, 1, int64(1)) // fail
//	NotEqual(t, 1, uint64(1)) // fail
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
