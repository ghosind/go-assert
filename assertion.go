package assert

import (
	"testing"
	"time"
)

type Assertion struct {
	t *testing.T
}

// New returns an assertion instance for verifying invariants.
func New(t *testing.T) *Assertion {
	a := new(Assertion)

	if t == nil {
		panic("parameter t is required")
	}
	a.t = t

	return a
}

// #########################
// ## Assertion Functions ##
// #########################

// DeepEqual tests deeply equality between actual and expect parameters.
func (a *Assertion) DeepEqual(actual, expect any, message ...string) error {
	return DeepEqual(a.t, actual, expect, message...)
}

// NotDeepEqual tests deeply inequality between actual and expected parameters.
func (a *Assertion) NotDeepEqual(actual, expect any, message ...string) error {
	return NotDeepEqual(a.t, actual, expect, message...)
}

// Nil tests whether a value is nil or not, and it'll fail when the value is not nil. It will
// always return false if the value is a bool, an integer, a floating number, a complex, or a
// string.
func (a *Assertion) Nil(val any, message ...string) error {
	return Nil(a.t, val, message...)
}

// NotNil tests whether a value is nil or not, and it'll fail when the value is nil. It will
// always return true if the value is a bool, an integer, a floating number, a complex, or a
// string.
func (a *Assertion) NotNil(val any, message ...string) error {
	return NotNil(a.t, val, message...)
}

// Panic expects the function fn to panic.
func (a *Assertion) Panic(fn func(), message ...string) (err error) {
	return Panic(a.t, fn, message...)
}

// NotPanic asserts that the function fn does not panic.
func (a *Assertion) NotPanic(fn func(), message ...string) (err error) {
	return NotPanic(a.t, fn, message...)
}

// ##########################
// ## Delegation Functions ##
// ##########################

// Cleanup registers a function to be called when the test (or subtest) and all its subtests
// complete. Cleanup functions will be called in last added, first called order.
func (a *Assertion) Cleanup(f func()) {
	a.t.Cleanup(f)
}

// Deadline reports the time at which the test binary will have exceeded the timeout specified by
// the -timeout flag.
//
// The ok result is false if the -timeout flag indicates “no timeout” (0).
func (a *Assertion) Deadline() (deadline time.Time, ok bool) {
	return a.t.Deadline()
}

// Error is equivalent to Log followed by Fail.
func (a *Assertion) Error(args ...any) {
	a.t.Error(args...)
}

// Errorf is equivalent to Logf followed by Fail.
func (a *Assertion) Errorf(format string, args ...any) {
	a.t.Errorf(format, args...)
}

// Fail marks the function as having failed but continues execution.
func (a *Assertion) Fail() {
	a.t.Fail()
}

// FailNow marks the function as having failed and stops its execution by calling runtime.Goexit
// (which then runs all deferred calls in the current goroutine). Execution will continue at the
// next test or benchmark. FailNow must be called from the goroutine running the test or benchmark
// function, not from other goroutines created during the test. Calling FailNow does not stop
// those other goroutines.
func (a *Assertion) FailNow() {
	a.t.FailNow()
}

// Failed reports whether the function has failed.
func (a *Assertion) Failed() bool {
	return a.t.Failed()
}

// Fatal is equivalent to Log followed by FailNow.
func (a *Assertion) Fatal(args ...any) {
	a.t.Fatal(args...)
}

// Fatalf is equivalent to Logf followed by FailNow.
func (a *Assertion) Fatalf(format string, args ...any) {
	a.t.Fatalf(format, args...)
}

// Helper marks the calling function as a test helper function. When printing file and line
// information, that function will be skipped. Helper may be called simultaneously from multiple
// goroutines.
func (a *Assertion) Helper() {
	a.t.Helper()
}

// Log formats its arguments using default formatting, analogous to Println, and records the text
// in the error log. For tests, the text will be printed only if the test fails or the -test.v
// flag is set. For benchmarks, the text is always printed to avoid having performance depend on
// the value of the -test.v flag.
func (a *Assertion) Log(args ...any) {
	a.t.Log(args...)
}

// Logf formats its arguments according to the format, analogous to Printf, and records the text in
// the error log. A final newline is added if not provided. For tests, the text will be printed
// only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to
// avoid having performance depend on the value of the -test.v flag.
func (a *Assertion) Logf(format string, args ...any) {
	a.t.Logf(format, args...)
}

// Name returns the name of the running (sub-) test or benchmark.
//
// The name will include the name of the test along with the names of any nested sub-tests. If two
// sibling sub-tests have the same name, Name will append a suffix to guarantee the returned name
// is unique.
func (a *Assertion) Name() string {
	return a.t.Name()
}

// Parallel signals that this test is to be run in parallel with (and only with) other parallel
// tests. When a test is run multiple times due to use of -test.count or -test.cpu, multiple
// instances of a single test never run in parallel with each other.
func (a *Assertion) Parallel() {
	a.t.Parallel()
}

// Run runs f as a subtest of t called name. It runs f in a separate goroutine and blocks until f
// returns or calls t.Parallel to become a parallel test. Run reports whether f succeeded (or at
// least did not fail before calling t.Parallel).
//
// Run may be called simultaneously from multiple goroutines, but all such calls must return before
// the outer test function for t returns.
func (a *Assertion) Run(name string, f func(*testing.T)) bool {
	return a.t.Run(name, f)
}

// Setenv calls os.Setenv(key, value) and uses Cleanup to restore the environment variable to its
// original value after the test.
//
// Because Setenv affects the whole process, it cannot be used in parallel tests or tests with
// parallel ancestors.
func (a *Assertion) Setenv(key, value string) {
	a.t.Setenv(key, value)
}

// Skip is equivalent to Log followed by SkipNow.
func (a *Assertion) Skip(args ...any) {
	a.t.Skip(args...)
}

// SkipNow marks the test as having been skipped and stops its execution by calling runtime.Goexit.
// If a test fails (see Error, Errorf, Fail) and is then skipped, it is still considered to have
// failed. Execution will continue at the next test or benchmark. See also FailNow. SkipNow must be
// called from the goroutine running the test, not from other goroutines created during the test.
// Calling SkipNow does not stop those other goroutines.
func (a *Assertion) SkipNow() {
	a.t.SkipNow()
}

// Skipf is equivalent to Logf followed by SkipNow.
func (a *Assertion) Skipf(format string, args ...any) {
	a.t.Skipf(format, args...)
}

// Skipped reports whether the test was skipped.
func (a *Assertion) Skipped() bool {
	return a.t.Skipped()
}

// TempDir returns a temporary directory for the test to use. The directory is automatically
// removed by Cleanup when the test and all its subtests complete. Each subsequent call to
// t.TempDir returns a unique directory; if the directory creation fails, TempDir terminates the
// test by calling Fatal.
func (a *Assertion) TempDir() string {
	return a.t.TempDir()
}
