package assert

import (
	"regexp"
	"testing"
)

// Match tests whether the string matches the regular expression or not.
func (a *Assertion) Match(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryMatchRegexp(a.T, false, val, pattern, message...)
}

// MatchNow tests whether the string matches the regular expression or not, and it will terminate
// the execution if it does not match.
func (a *Assertion) MatchNow(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryMatchRegexp(a.T, true, val, pattern, message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not. It will panic if the pattern is not a valid regular expression.
func (a *Assertion) MatchString(val, pattern string, message ...string) error {
	a.Helper()

	regPattern := regexp.MustCompile(pattern)

	return tryMatchRegexp(a.T, false, val, regPattern, message...)
}

// MatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not. It will terminate the execution if it does not match, and it will panic if
// the pattern is not a valid regular expression.
func (a *Assertion) MatchStringNow(val, pattern string, message ...string) error {
	a.Helper()

	regPattern := regexp.MustCompile(pattern)

	return tryMatchRegexp(a.T, true, val, regPattern, message...)
}

// NotMatch tests whether the string matches the regular expression or not, and it set the result
// to fail if the string matches the pattern.
func (a *Assertion) NotMatch(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, false, val, pattern, message...)
}

// NotMatchNow tests whether the string matches the regular expression or not, and it will
// terminate the execution if the string matches the pattern.
func (a *Assertion) NotMatchNow(val string, pattern *regexp.Regexp, message ...string) error {
	a.Helper()

	return tryNotMatchRegexp(a.T, true, val, pattern, message...)
}

// MatchString will compile the pattern and test whether the string matches the regular expression
// or not, and it set the result to fail if the string matches the pattern. It will also panic if
// the pattern is not a valid regular expression.
func (a *Assertion) NotMatchString(val, pattern string, message ...string) error {
	a.Helper()

	regPattern := regexp.MustCompile(pattern)

	return tryNotMatchRegexp(a.T, false, val, regPattern, message...)
}

// NotMatchStringNow will compile the pattern and test whether the string matches the regular
// expression or not, and it set the result to fail if the string matches the pattern. It will
// terminate the execution if the string matches the pattern, and it will panic if the pattern is
// not a valid regular expression.
func (a *Assertion) NotMatchStringNow(val, pattern string, message ...string) error {
	a.Helper()

	regPattern := regexp.MustCompile(pattern)

	return tryNotMatchRegexp(a.T, true, val, regPattern, message...)
}

// tryMatchRegexp tries to test whether the string matches the regular expression pattern or not,
// and it'll fail if the string does not match.
func tryMatchRegexp(
	t *testing.T,
	failedNow bool,
	val string,
	pattern *regexp.Regexp,
	message ...string,
) error {
	t.Helper()

	if pattern.Match([]byte(val)) {
		return nil
	}

	err := newAssertionError("The input did not match the regular expression", message...)
	failed(t, err, failedNow)
	return err
}

// tryNotMatchRegexp tries to test whether the string matches the regular expression pattern or
// not, and it'll fail if the string matches the pattern.
func tryNotMatchRegexp(
	t *testing.T,
	failedNow bool,
	val string,
	pattern *regexp.Regexp,
	message ...string,
) error {
	t.Helper()

	if !pattern.Match([]byte(val)) {
		return nil
	}

	err := newAssertionError("The input match the regular expression", message...)
	failed(t, err, failedNow)
	return err
}
