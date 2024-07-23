package assert

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// ContainsString tests whether the string contains the substring or not, and it set the result to
// fail if the string does not contains the substring.
//
//	a := assert.New(t)
//	a.ContainsString("Hello world", "") // success
//	a.ContainsString("Hello world", "Hello") // success
//	a.ContainsString("Hello world", "world") // success
//	a.ContainsString("Hello world", "hello") // fail
func (a *Assertion) ContainsString(str, substr string, message ...any) error {
	a.Helper()

	return tryContainsString(a.T, false, str, substr, message...)
}

// ContainsStringNow tests whether the string contains the substring or not, and it will terminate the
// execution if the string does not contains the substring.
//
//	a := assert.New(t)
//	a.ContainsStringNow("Hello world", "") // success
//	a.ContainsStringNow("Hello world", "Hello") // success
//	a.ContainsStringNow("Hello world", "world") // success
//	a.ContainsStringNow("Hello world", "hello") // fail and stop the execution
//	// never runs
func (a *Assertion) ContainsStringNow(str, substr string, message ...any) error {
	a.Helper()

	return tryContainsString(a.T, true, str, substr, message...)
}

// NotContainsString tests whether the string contains the substring or not, and it set the result
// to fail if the string contains the substring.
//
//	a := assert.New(t)
//	a.NotContainsString("Hello world", "hello") // success
//	a.NotContainsString("Hello world", "") // fail
//	a.NotContainsString("Hello world", "Hello") // fail
//	a.NotContainsString("Hello world", "world") // fail
func (a *Assertion) NotContainsString(str, substr string, message ...any) error {
	a.Helper()

	return tryNotContainsString(a.T, false, str, substr, message...)
}

// NotContainsStringNow tests whether the string contains the substring or not, and it will terminate the
// execution if the string contains the substring.
//
//	a := assert.New(t)
//	a.NotContainsStringNow("Hello world", "hello") // success
//	a.NotContainsStringNow("Hello world", "Hello") // fail and stop the execution
//	// never runs
func (a *Assertion) NotContainsStringNow(str, substr string, message ...any) error {
	a.Helper()

	return tryNotContainsString(a.T, true, str, substr, message...)
}

// tryContainsString tries to test whether the string contains the substring or not, and it'll
// fail if the string does not contains the substring.
func tryContainsString(
	t *testing.T,
	failedNow bool,
	str, substr string,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return strings.Contains(str, substr) },
		failedNow,
		fmt.Sprintf(defaultErrMessageContainsString, substr),
		message...,
	)
}

// tryNotContainsString tries to test whether the string contains the substring or not, and it'll
// fail if the string contains the substring.
func tryNotContainsString(
	t *testing.T,
	failedNow bool,
	str, substr string,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return !strings.Contains(str, substr) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotContainsString, substr),
		message...,
	)
}

// HasPrefixString tests whether the string has the prefix string or not, and it set the result to
// fail if the string does not have the prefix string.
//
//	a := assert.New(t)
//	a.HasPrefixString("Hello world", "") // success
//	a.HasPrefixString("Hello world", "Hello") // success
//	a.HasPrefixString("Hello world", "world") // fail
//	a.HasPrefixString("Hello world", "hello") // fail
func (a *Assertion) HasPrefixString(str, prefix string, message ...any) error {
	a.Helper()

	return tryHasPrefixString(a.T, false, str, prefix, message...)
}

// HasPrefixStringNow tests whether the string has the prefix string or not, and it will terminate
// the execution if the string does not have the prefix string.
//
//	a := assert.New(t)
//	a.HasPrefixStringNow("Hello world", "") // success
//	a.HasPrefixStringNow("Hello world", "Hello") // success
//	a.HasPrefixStringNow("Hello world", "hello") // fail and stop the execution
//	// never runs
func (a *Assertion) HasPrefixStringNow(str, prefix string, message ...any) error {
	a.Helper()

	return tryHasPrefixString(a.T, true, str, prefix, message...)
}

// NotHasPrefixString tests whether the string has the prefix string or not, and it set the result
// to fail if the string have the prefix string.
//
//	a := assert.New(t)
//	a.NotHasPrefixString("Hello world", "hello") // success
//	a.NotHasPrefixString("Hello world", "world") // success
//	a.NotHasPrefixString("Hello world", "") // fail
//	a.NotHasPrefixString("Hello world", "Hello") // fail
func (a *Assertion) NotHasPrefixString(str, prefix string, message ...any) error {
	a.Helper()

	return tryNotHasPrefixString(a.T, false, str, prefix, message...)
}

// NotHasPrefixStringNow tests whether the string has the prefix string or not, and it will
// terminate the execution if the string have the prefix string.
//
//	a := assert.New(t)
//	a.NotHasPrefixStringNow("Hello world", "hello") // success
//	a.NotHasPrefixStringNow("Hello world", "world") // success
//	a.NotHasPrefixStringNow("Hello world", "Hello") // fail and stop the execution
//	// never runs
func (a *Assertion) NotHasPrefixStringNow(str, prefix string, message ...any) error {
	a.Helper()

	return tryNotHasPrefixString(a.T, true, str, prefix, message...)
}

// tryHasPrefixString tries to test whether the string has the prefix string or not, and it'll fail
// if the string does not have the prefix string.
func tryHasPrefixString(
	t *testing.T,
	failedNow bool,
	str, prefix string,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return strings.HasPrefix(str, prefix) },
		failedNow,
		fmt.Sprintf(defaultErrMessageHasPrefixString, prefix),
		message...,
	)
}

// tryNotHasPrefixString tries to test whether the string has the prefix string or not, and it'll
// fail if the string has the prefix string.
func tryNotHasPrefixString(
	t *testing.T,
	failedNow bool,
	str, prefix string,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return !strings.HasPrefix(str, prefix) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotHasPrefixString, prefix),
		message...,
	)
}

// HasSuffixString tests whether the string has the suffix string or not, and it set the result to
// fail if the string does not have the suffix string.
//
//	a := assert.New(t)
//	a.HasSuffixString("Hello world", "") // success
//	a.HasSuffixString("Hello world", "world") // success
//	a.HasSuffixString("Hello world", "World") // fail
//	a.HasSuffixString("Hello world", "hello") // fail
func (a *Assertion) HasSuffixString(str, suffix string, message ...any) error {
	a.Helper()

	return tryHasSuffixString(a.T, false, str, suffix, message...)
}

// HasSuffixStringNow tests whether the string has the suffix string or not, and it will terminate
// the execution if the string does not have the suffix string.
//
//	a := assert.New(t)
//	a.HasSuffixStringNow("Hello world", "") // success
//	a.HasSuffixStringNow("Hello world", "world") // success
//	a.HasSuffixStringNow("Hello world", "World") // fail and stop the execution
//	// never runs
func (a *Assertion) HasSuffixStringNow(str, suffix string, message ...any) error {
	a.Helper()

	return tryHasSuffixString(a.T, true, str, suffix, message...)
}

// NotHasSuffixString tests whether the string has the suffix string or not, and it set the result
// to fail if the string have the suffix string.
//
//	a := assert.New(t)
//	a.NotHasSuffixString("Hello world", "Hello") // success
//	a.NotHasSuffixString("Hello world", "World") // success
//	a.NotHasSuffixString("Hello world", "") // fail
//	a.NotHasSuffixString("Hello world", "world") // fail
func (a *Assertion) NotHasSuffixString(str, suffix string, message ...any) error {
	a.Helper()

	return tryNotHasSuffixString(a.T, false, str, suffix, message...)
}

// NotHasSuffixStringNow tests whether the string has the suffix string or not, and it will
// terminate the execution if the string have the suffix string.
//
//	a := assert.New(t)
//	a.NotHasSuffixStringNow("Hello world", "hello") // success
//	a.NotHasSuffixStringNow("Hello world", "World") // success
//	a.NotHasSuffixStringNow("Hello world", "world") // fail and stop the execution
//	// never runs
func (a *Assertion) NotHasSuffixStringNow(str, suffix string, message ...any) error {
	a.Helper()

	return tryNotHasSuffixString(a.T, true, str, suffix, message...)
}

// tryHasSuffixString tries to test whether the string has the suffix string or not, and it'll fail
// if the string does not have the suffix string.
func tryHasSuffixString(
	t *testing.T,
	failedNow bool,
	str, suffix string,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return strings.HasSuffix(str, suffix) },
		failedNow,
		fmt.Sprintf(defaultErrMessageHasSuffixString, suffix),
		message...,
	)
}

// tryNotHasSuffixString tries to test whether the string has the suffix string or not, and it'll
// fail if the string has the suffix string.
func tryNotHasSuffixString(
	t *testing.T,
	failedNow bool,
	str, suffix string,
	message ...any,
) error {
	t.Helper()

	return test(
		t,
		func() bool { return !strings.HasSuffix(str, suffix) },
		failedNow,
		fmt.Sprintf(defaultErrMessageNotHasSuffixString, suffix),
		message...,
	)
}

// Match tests whether the string matches the regular expression or not.
//
//	a := assert.New(t)
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	a.Match("http://example.com", pattern) // success
//	a.Match("example.com", pattern) // fail
func (a *Assertion) Match(val string, pattern *regexp.Regexp, message ...any) error {
	a.Helper()

	return tryMatchRegexp(a.T, false, val, pattern, "", message...)
}

// MatchNow tests whether the string matches the regular expression or not, and it will terminate
// the execution if it does not match.
//
//	a := assert.New(t)
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	a.MatchNow("http://example.com", pattern) // success
//	a.MatchNow("example.com", pattern) // fail and terminate
//	// never run
func (a *Assertion) MatchNow(val string, pattern *regexp.Regexp, message ...any) error {
	a.Helper()

	return tryMatchRegexp(a.T, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not. It will panic if the pattern is not a valid regular expression.
//
//	a := assert.New(t)
//	a.MatchString("http://example.com", `^https?:\/\/`) // success
//	a.MatchString("example.com", `^https?:\/\/`) // fail
func (a *Assertion) MatchString(val, pattern string, message ...any) error {
	a.Helper()

	return tryMatchRegexp(a.T, false, val, nil, pattern, message...)
}

// MatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not. It will terminate the execution if it does not match, and it will panic if
// the pattern is not a valid regular expression.
//
//	a := assert.New(t)
//	a.MatchStringNow("http://example.com", `^https?:\/\/`) // success
//	a.MatchStringNow("example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func (a *Assertion) MatchStringNow(val, pattern string, message ...any) error {
	a.Helper()

	return tryMatchRegexp(a.T, true, val, nil, pattern, message...)
}

// NotMatch tests whether the string matches the regular expression or not, and it set the result
// to fail if the string matches the pattern.
//
//	a := assert.New(t)
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	a.NotMatch("example.com", pattern) // success
//	a.NotMatch("http://example.com", pattern) // fail
func (a *Assertion) NotMatch(val string, pattern *regexp.Regexp, message ...any) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, false, val, pattern, "", message...)
}

// NotMatchNow tests whether the string matches the regular expression or not, and it will
// terminate the execution if the string matches the pattern.
//
//	a := assert.New(t)
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	a.NotMatchNow("example.com", pattern) // success
//	a.NotMatchNow("http://example.com", pattern) // fail and terminate
//	// never run
func (a *Assertion) NotMatchNow(val string, pattern *regexp.Regexp, message ...any) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not, and it set the result to fail if the string matches the pattern. It will also panic if
// the pattern is not a valid regular expression.
//
//	a := assert.New(t)
//	a.NotMatchString("example.com", `^https?:\/\/`) // success
//	a.NotMatchString("http://example.com", `^https?:\/\/`) // fail
func (a *Assertion) NotMatchString(val, pattern string, message ...any) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, false, val, nil, pattern, message...)
}

// NotMatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not, and it set the result to fail if the string matches the pattern. It will
// terminate the execution if the string matches the pattern, and it will panic if the pattern is
// not a valid regular expression.
//
//	a := assert.New(t)
//	a.NotMatchStringNow("example.com", `^https?:\/\/`) // success
//	a.NotMatchStringNow("http://example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func (a *Assertion) NotMatchStringNow(val, pattern string, message ...any) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, true, val, nil, pattern, message...)
}

// tryMatchRegexp tries to test whether the string matches the regular expression pattern or not,
// and it'll fail if the string does not match.
func tryMatchRegexp(
	t *testing.T,
	failedNow bool,
	val string,
	pattern *regexp.Regexp,
	patternStr string,
	message ...any,
) error {
	t.Helper()

	if pattern == nil {
		pattern = regexp.MustCompile(patternStr)
	}

	return test(
		t,
		func() bool { return pattern.Match([]byte(val)) },
		failedNow,
		defaultErrMessageMatch,
		message...,
	)
}

// tryNotMatchRegexp tries to test whether the string matches the regular expression pattern or
// not, and it'll fail if the string matches the pattern.
func tryNotMatchRegexp(
	t *testing.T,
	failedNow bool,
	val string,
	pattern *regexp.Regexp,
	patternStr string,
	message ...any,
) error {
	t.Helper()

	if pattern == nil {
		pattern = regexp.MustCompile(patternStr)
	}

	return test(
		t,
		func() bool { return !pattern.Match([]byte(val)) },
		failedNow,
		defaultErrMessageNotMatch,
		message...,
	)
}
