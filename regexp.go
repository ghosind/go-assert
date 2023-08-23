package assert

import (
	"regexp"
	"testing"
)

// Match tests whether the string matches the regular expression or not.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assertion.Match("http://example.com", pattern) // success
//	assertion.Match("example.com", pattern) // fail
func (a *Assertion) Match(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryMatchRegexp(a.T, false, val, pattern, "", message...)
}

// MatchNow tests whether the string matches the regular expression or not, and it will terminate
// the execution if it does not match.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assertion.MatchNow("http://example.com", pattern) // success
//	assertion.MatchNow("example.com", pattern) // fail and terminate
//	// never run
func (a *Assertion) MatchNow(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryMatchRegexp(a.T, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not. It will panic if the pattern is not a valid regular expression.
//
//	assertion.MatchString("http://example.com", `^https?:\/\/`) // success
//	assertion.MatchString("example.com", `^https?:\/\/`) // fail
func (a *Assertion) MatchString(val, pattern string, message ...string) error {
	a.Helper()

	return tryMatchRegexp(a.T, false, val, nil, pattern, message...)
}

// MatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not. It will terminate the execution if it does not match, and it will panic if
// the pattern is not a valid regular expression.
//
//	assertion.MatchStringNow("http://example.com", `^https?:\/\/`) // success
//	assertion.MatchStringNow("example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func (a *Assertion) MatchStringNow(val, pattern string, message ...string) error {
	a.Helper()

	return tryMatchRegexp(a.T, true, val, nil, pattern, message...)
}

// NotMatch tests whether the string matches the regular expression or not, and it set the result
// to fail if the string matches the pattern.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assertion.NotMatch("example.com", pattern) // success
//	assertion.NotMatch("http://example.com", pattern) // fail
func (a *Assertion) NotMatch(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, false, val, pattern, "", message...)
}

// NotMatchNow tests whether the string matches the regular expression or not, and it will
// terminate the execution if the string matches the pattern.
//
//	pattern := regexp.MustCompile(`^https?:\/\/`)
//	assertion.NotMatchNow("example.com", pattern) // success
//	assertion.NotMatchNow("http://example.com", pattern) // fail and terminate
//	// never run
func (a *Assertion) NotMatchNow(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, true, val, pattern, "", message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not, and it set the result to fail if the string matches the pattern. It will also panic if
// the pattern is not a valid regular expression.
//
//	assertion.NotMatchString("example.com", `^https?:\/\/`) // success
//	assertion.NotMatchString("http://example.com", `^https?:\/\/`) // fail
func (a *Assertion) NotMatchString(val, pattern string, message ...string) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, false, val, nil, pattern, message...)
}

// NotMatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not, and it set the result to fail if the string matches the pattern. It will
// terminate the execution if the string matches the pattern, and it will panic if the pattern is
// not a valid regular expression.
//
//	assertion.NotMatchStringNow("example.com", `^https?:\/\/`) // success
//	assertion.NotMatchStringNow("http://example.com", `^https?:\/\/`) // fail and terminate
//	// never run
func (a *Assertion) NotMatchStringNow(val, pattern string, message ...string) error {
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
	message ...string,
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
	message ...string,
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
